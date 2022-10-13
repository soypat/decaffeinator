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