# Appending to a slice
It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append.
```go
func append(s []T, values ...T) []T
```
The first parameter s of append is a slice of type `T`, and the rest are `T` values to append to the slice.

The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

(To learn more about slices, read the [Slices: usage and internals](https://go.dev/blog/slices-intro) article.)

## Appending a slice to a slice
Go has an easy way of doing this using the append built-in and varargs

```go
s := []int{1, 2, 3, 4}
sn := []int{5, 6, 7}
joined := append(s, sn...)
// joined is now [1, 2, 3, 4, 5, 6, 7]
```
The `...` operator "unpacks" the slice (strictly speaking no unpacking is done).


## Pitfalls
Go's append mechanism is surprisingly simple underneath the hood:
- Appending adds elements to the end of the slice and returns a new slice with old elements and a new length.
- If there is not enough capacity in the slice a new slice is allocated with double the size and contents copied and appended to it.

Although the rules regarding slices are simple the fact that there are pointers betwixt sometimes creates unintuitive results.

### Overwriting 
The example below overwrites elements appended in an earlier call since append is called on the same slice twice.
```go
s := []int{1, 2}
s1 := append(s, 3)
s2 := append(s, 4)
fmt.Println(s1, s2)
```
<details><summary>Result</summary><pre>
1 2 4
1 2 4
</pre></details>

**Solution:** You are most likely thinking of slices as *values* if this is an issue, this will bring many misconceptions in the long run of using Go. If you are using the `append` functionality in your program it is because the slice you are working with is an *accumulator* kind of variable. Declare one slice for each kind of data you need to accumulate and put corresponding data in each of the slices with `append` always using the usual pattern `slice = append(slice, data)`.


### Memory leaks
Advancing the slice start will not free up the memory in the front of the slice to be available for garbage collection. There are some times you'd want to do this, maybe during processing of entities when you know you'll
dispose of the slice when done. 

```go
// The memory at s[:i] is now not part of the slice used memory. 
s = s[i:]
```

**Solution:** Copy-to-front pattern. Copy is quite fast on modern systems. Note that aliased copy in Go is always defined behaviour, unlike other UB-heavy languages.
```go
// Equivalent to s = s[:i] but uses slice memory efficiently.
n := copy(s[:0], s[i:])
s = s[:n]
```

### Loose references
Be wary of taking pointers from a slice you are appending to. If the slice is grown in capacity all data will be moved and new data appended to new slice.

```go
ptr := &s[i]
// This may generate a new slice and now ptr does not point to s!
s = append(s, data) 
newPtr := &s[i] // May not be the same as ptr!
```

**Solution:** Don't keep long lived pointers to a slice that will be appended to. This goes back to the saying: Know your memory. Know the lifetime of your data structures. It is fine to take a pointer to a slice for a calculation, but plan on not keeping that pointer if you plan on appending to it.

### Heap allocations
The garbage collector is a wonderful thing, but be wary if you plan on writing high performance code. There are ways to get around heap allocations by checking slice capacity. These algorithms are called
"heapless" since they avoid using the heap when possible.

```go
newElements := len(elements)
free := cap(s)-len(s)
if free >= newElements {
    // This append call will never allocate since we checked free space.
    s = append(s, elements...)
} else {
    println("not enough space!")
}
```

