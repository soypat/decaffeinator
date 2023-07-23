# Inline functions
Go's functions are what is known as "First class citizens". This is just a fancy way
of saying functions are also values in Go and can be treated the same way as integers and strings.

In Go one can define a function within the context of a function and assign it to a variable.
The example below defines a new function and assign it to `add` within `main`.
The program will print out `3` followed by `6` on separate lines.

```go
package main

import "fmt"

func main() {
    add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(1, 2)) // 3

	fmt.Println(add(4, 2)) // 6
}
```
Note that although the function is assigned to a variable called `add` the function itself
does **not** have a name! The compiler sees it as `main.func1`. This is why you'll also
see these functions called "anonymous functions" in literature.


If the `add` function were defined outside it would have the same effect:

```go
package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func main() {

	fmt.Println(add(1, 2)) // 3

	fmt.Println(add(4, 2)) // 6
}
```

Why is it then that Go lets us define functions inline if we can _seemingly_
omit this flashy feature?

### When to use inline functions
#### Capturing local variables or Closures
When done we call the resulting inline function a **closure**.
Closures are a powerful construct when combined with function modifier keywords `defer` and `go`.

Below is a program that captures the `a` integer variable and adds to it on every call.

```go
a := 0
adder := func(b int) int {
    a += b
    return a
}

fmt.Println(adder(2)) // 2
fmt.Println(adder(3)) // 5
fmt.Println(a)        // 5
```
The above behaviour can be mind boggling at first, since it's not immediately obvious from what
we've learned how `a` is being modified if it's not an argument to the inline function.

The reason behind it is because when `a` is used inside of the inline function
a is received by the closure as a reference (pointer).
One may ask how is that possible if `a` is a non-pointer type; that's because of how 
inline functions work: they "capture" variable by reference. Any variable external to the inline functions scope that one uses within the inline function will be linked to the original variable.
To avoid this behaviour pass the variable as an argument to the inline function.


#### Locality
Inline functions are anonymous and can't be accessed externally or even from within 
other functions within the same package. This has strong security implications.

There is also the benefit of having logic close to where its used. Since Go functions are first class citizens they can be arguments to a function. These functions are called _higher order functions_.

```go
// SuperRandom returns a super random number using a 
// not so random source function "normalRandom".
// (this is not actually more random!)
func SuperRandom(normalRandom func() int) int {
    superrand := 12345678
    for i := 0; i < 3; i++ {
        rand := normalRandom()
        superrand = superrand*7 + rand*31
    }
    return superrand
}
```
The `SuperRandom` function above expects a `func() int` as an argument (a function with no parameters that returns an integer). When calling the function above we could define a inline
function to satisfy it's parameter:

```go
a := 287117
notSoRand := func() int {
    a = a * 7
    return a
}
superrand1 := SuperRandom(notSoRand)
superrand2 := SuperRandom(notSoRand)
superrand3 := SuperRandom(notSoRand)

fmt.Println(superrand1, superrand2, superrand3)
```

By locally defining a function to pass as a value to another function
we make the code more readable than if the function were defined far away.
