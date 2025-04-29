# Pointers
You may have or not heard of pointers, the famous core concept of many programming languages.

## What are they, from a functional perspective
Pointers allow a calling function access to memory outside of what is defined in its stack frame- so variables outside of the scope that have been defined in the function. In C-like languages pointer types are expressed with an asterisk around the type the pointer would reference:

```go
// add2 receives a pointer that references an integer.
func add2(p *int)
```

The function `add2` is free to read or write to the integer at p's address. 

To read the address at 

One reads from a pointer by "dereferencing" the address with the asterisk operator.

```go
var value int = *p
```

Similarly, we may assign to the address by dereferencing the pointer in the same way we read from it:
```go
*p = 2
```

So, the function which adds 2 to an integer given the integer's address could look like so:

```go
func add2(p *int) {
    currentValue := *p
    newValue := currentValue+2
    *p = newValue
}
```

If the price of a line of code is unnafordable then we can use the plus-equal operator:

```go
func add2(p *int) {
    *p += 2
}
```

## Caveats

Here are some myths you may have heard regarding pointers:
- **Myth**: Pointers are faster
- **Myth**: Pointers are slower
- **Myth**: Use pointers for structs
- **Myth**: Use pointers when you want functions to modify the arguments

Computers are extremely complex systems, there is not one-fits-all rule for using pointers. This is not to discourage the use of pointers, by all means **use them, and use them alot**! This is the best way to learn how they work. But if you are building something complex and can do without pointers, consider omitting their use to avoid bugs that can come along with the added complexity of pointers.

There is one clear-cut use case for pointers in Go and that is to implement interfaces that require modifying the receiver's data. More on that in the interfaces section.
- **Fact**: Use pointer receivers when implementing interfaces which require modifying the receiver's data.