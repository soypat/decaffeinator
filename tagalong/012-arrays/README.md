# Arrays
The type `[n]T` is an array of *n* values of type `T`. Arrays
are mutable.

The expression
```go
var a [10]int
```
declares a variable a as an array of ten integers and the expression
```go
var z = [...]string{"A", "B", "Z"}
```
declares and initializes an array of 3 strings with shown values. The `...` auto deduces its length.

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays...