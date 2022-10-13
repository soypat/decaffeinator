# Slices
An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type `[]T` is a slice with elements of type `T`.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

```
a[low : high]
```

This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a:

```
a[1:4]
```

## Slice defaults
When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.

For the array
```go
var a [10]int
```
these slice expressions are equivalent:

```go
a[0:10]
a[:10]
a[0:]
a[:]
```

## Slice literals
A slice literal is like an array literal without the length.

This is an array literal:

```go
[3]bool{true, true, false}
```

And this creates the same array as above, then builds a slice that references it:
```go
slice := []bool{true, true, false}
// equivalent to
array := [3]bool{true, true, false}
slice := array[:]
```

## Zero value and nil slices
Nil slices
The zero value of a slice is `nil`. Similar to `null` and `None` in other languages.

A `nil` slice has a length and capacity of 0 and has no underlying array.

## Creating a slice with make
Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array:
```go
a := make([]int, 5)  // len(a)=5
```

**Advanced topic** Capacity: To specify a capacity, pass a third argument to make:
```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

## Slice of slices
Slices can contain any type, including other slices.
```go
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
    // The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
    // print board:
    for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
```

