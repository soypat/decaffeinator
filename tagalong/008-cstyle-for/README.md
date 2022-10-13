# For
Go has only one looping construct, the for loop.

The basic for loop has three components separated by semicolons:

1. the **init statement**: executed before the first iteration
2. the **condition expression**: evaluated before every iteration
3. the **post statement**: executed at the end of every iteration

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

The loop will stop iterating once the boolean condition evaluates to false.

Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.

## While
Worth mentioning the post and init statements are optional

```go
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
```
At that point you can drop the semicolons: C's while is spelled for in Go.
```go
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
```
## Forever
If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
```go
	for {
	}
```