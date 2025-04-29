# Pointers and Slices
In Go slices are "fat pointers", which is to say they are not a pointer in itself but rather a struct that contains a pointer and length data.

This makes them less pointery than equivalent data structures in other languages such as Python's `List` type which allow a function to modify the caller's list from inside the function.

Due to the "fat" nature of Go's slices it is common for a function that modifies a slice's length to return the resulting slice, just like the `append` function. This is a widespread pattern in the Go ecosystem as well as in the standard library. See the [`strconv.Append`](https://pkg.go.dev/strconv#AppendInt) family of functions as well as the [`fmt.Appendf`](https://pkg.go.dev/fmt#Appendf) function.

