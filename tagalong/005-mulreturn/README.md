# Functions (continued)

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

One can also omit returned value type for same consecutive returned values such
that the following two function signatures are functionally identical:

```go
func collatz(a int) (down int, up int)

func collatz(a int) (down, up int)
```