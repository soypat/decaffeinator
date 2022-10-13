# If
Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.

Like for, the if statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if.

```go
    if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
    v++ // compile time error.
```