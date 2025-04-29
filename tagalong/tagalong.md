# Tagalong - Introduction
This document aims to help Pythonistas understand Go syntax at a glance by
leveraging Python and Go examples that do identical work and a short explanation
of Go's syntax preceding said examples.

Hopefully Pythonistas will learn Go is about the same level as Python when
one speaks of "high/low" level programming languages. 

There is a version of this document that also includes a Zig program 
to drive the previous point home and for those curious on a lower level language
that could replace C in academia.

**You'll find the examples which the section will refer to at the end of the section.**

*WIP*: This document was generated programatically.
Find the source code at [github.com/soypat/decaffeinator](https://github.com/soypat/decaffeinator).
# Hello World
No explanation here. This is just the common fiat language demo.
### Python (hello)
```python
print("Hello, world!")
```
### Go (hello)
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

```
**Output**:
```plaintext
Hello, world!
```

# Packages

Every Go program is made up of packages.

Programs start running in package `main`.

This program is using the packages with import paths "fmt" and "math/rand".

By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package `rand`.

Note: The environment in which these programs are executed is deterministic, so each time you run the example program rand.Intn will return the same number.

(To see a different number, seed the number generator; see rand.Seed. Time is constant in the playground, so you will need to use something else as the seed.)

## Import statement
This code groups the imports into a parenthesized, "factored" import statement.

## Exported names
In Go, a name is exported if it begins with a capital letter. For example, `Pizza` is an exported name, as is `Pi`, which is exported from the `math` package.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.
### Python (packages)
```python
import math
import random

print("my favorite number is",random.randint(0, 10),"and", math.pi)
```
### Go (packages)
```go
package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10), "and", math.Pi)
}

```
**Output**:
```plaintext
My favorite number is 4 and 3.141592653589793
```

# Functions
A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.

(For more about why types look the way they do, see the article on Go's declaration syntax.)

When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
Thus, the following function signatures are equivalent:

```go
func add(x int, y int) int

func add(x, y int) int
```
### Python (funcadd)
```python
def add(a:int, b:int) -> int:
    return a+b

print(add(42,13))
```
### Go (funcadd)
```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

```
**Output**:
```plaintext
55
```

# Functions (continued)

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

One can also omit returned value type for same consecutive returned values such
that the following two function signatures are functionally identical:

```go
func collatz(a int) (down int, up int)

func collatz(a int) (down, up int)
```
### Python (mulreturn)
```python
def collatz(a:int) -> tuple[int, int]:
    down = a//2
    up = 3*a+1
    return down, up

v = 60
down, up = collatz(v)
print(f"empezando en {v} hay que saber subir {up}, y bajar {down}")
```
### Go (mulreturn)
```go
package main

import "fmt"

func collatz(a int) (down int, up int) {
	down = a / 2
	up = a*3 + 1
	return down, up
}

func main() {
	const v = 60
	down, up := collatz(v)
	fmt.Printf("empezando en %d hay que saber subir %d, y bajar %d", v, up, down)
}

```
**Output**:
```plaintext
empezando en 60 hay que saber subir 181, y bajar 30
```

# Variables
The var statement declares a list of variables; as in function argument lists, the type is last.

Inside a function, the `:=` short assignment statement can be used in place of a var declaration with implicit type.

## Go's basic types

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

## Zero value
Variables declared without an explicit initial value are given their zero value.

The zero value is:

1. `0` for numeric types,
2. `false` for the boolean type, and
3. `""` (the empty string) for strings.

### Python (variables)
```python
s = ""
s2 = "Hello!"

i, j = 0, 42
k = 12
print(f"\"{s}\" \"{s2}\" {i} {j} {k}")
```
### Go (variables)
```go
package main

import "fmt"

func main() {
	var s string
	s2 := "Hello!"

	var i, j int
	j = 42
	k := 12
	fmt.Printf("%q %q %d %d %d", s, s2, i, j, k)
}

```
**Output**:
```plaintext
"" "Hello!" 0 42 12
```

# Variables (continued)
A var statement can be at package or function level. We see both in this example.

Outside a function, every statement begins with a keyword (`var`, `func`, and so on) and so the `:=` construct is not available.

### Python (variables2)
```python
c, python, java = False, False, False

s = 0
s2 = "This is long text"
start, stop = 1., 20.

afloat = 6.02
casted = int(afloat)
words = s2.split(sep=" ")
print(c, python, java, s, words, start, stop, afloat, casted)


```
### Go (variables2)
```go
package main

import (
	"fmt"
	"strings"
)

var c, python, java bool

var (
	s           int64
	s2          string  = "This is long text"
	start, stop float64 = 1, 20
)

func main() {
	afloat := 6.02
	casted := int(afloat)
	words := strings.Split(s2, " ")
	fmt.Println(c, python, java, s, words, start, stop, afloat, casted)
}

```
**Output**:
```plaintext
false false false 0 [This is long text] 1 20 6.02 6
```

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
### Python (cstyle-for)
```python
sum = 0
i = 0
while i < 10:
    sum += i
    i+=1
    print(sum)

```
### Go (cstyle-for)
```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
}

```
**Output**:
```plaintext
0
1
3
6
10
15
21
28
36
45
```

# Switch
A `switch` statement is a shorter way to write a sequence of `if` - `else` statements. It runs the first case whose value is equal to the condition expression.

Go's `switch` is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the `break` statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

**Switch cases evaluate cases from top to bottom, stopping when a case succeeds.**

For example:
```go
switch i {
case 0:
case f():
}
```
does not call `f` if `i==0`.

## Switch with no condition
Switch without a condition is the same as switch true.

This construct can be a clean way to write long if-then-else chains.

```go
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
```
### Python (switch)
```python
import datetime
import time

today = datetime.date.today()

# This control structure is poorly supported
# in VSCode as of October 2022, Python 3.10. Quite hard to write.
match today.weekday():
    case 5:
        print("Today.")
    case 4:
        print("Tomorrow.")
    case 5:
        print("In two days.")
    case _:
        print("Too far away.")
```
### Go (switch)
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch today {
	case time.Saturday:
		fmt.Println("Today.")
	case time.Friday:
		fmt.Println("Tomorrow.")
	case time.Thursday:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

```
**Output**:
```plaintext
When's Saturday?
Too far away.
```

# Structs
A `struct` is a collection of fields.

Struct fields are accessed using a dot.

## Struct Literals
A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the `Name: syntax`. (And the order of named fields is irrelevant.)


You can also declare structs without specifying the fields based on ordering, i.e.

```go
type Vertex struct {
	X int
	Y int
}

var v = Vertex{1, 2} // X=1, Y=2
```
### Python (struct)
```python
class Vertex:
    def __init__(self, x:int=0, y:int=0):
        self.x= int(x)
        self.y= int(y)

    def __str__(self):
        return f"Vertex({self.x}, {self.y})"

    def __setattr__(self, __name: str, __value: any) -> None:
        super().__setattr__(__name, int(__value))

v0 = Vertex()
print(v0)
a = 0

v1 = Vertex(1, 2)
print(v1.x)

# This is a bug with no __setattr__ method.
# I didn't look for a spurious case, this was taken as is from
# https://go.dev/tour/moretypes/4
v1.x = 1e9
print(v1)


```
### Go (struct)
```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v0 := Vertex{}
	fmt.Println(v0)

	v1 := Vertex{X: 1, Y: 2}
	fmt.Println("X:", v1.X)

	v1.X = 1e9
	fmt.Println("new v1:", v1)

	// You may also print the struct with it's fields with the +v formatting directive.
	fmt.Printf("%+v", v1)
}

```
**Output**:
```plaintext
{0 0}
X: 1
new v1: {1000000000 2}
{X:1000000000 Y:2}
```

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
### Python (arrays)
```python
# Static declaration of list (mutable)
a = list[str](["",""])
a[0] = "Hello"
a[1] = "World"
print(a[0], a[1])
print(a)

# Dynamic declaration of tuple (non-mutable)
primes = (2, 3, 5, 7, 11, 13)
print(primes)

```
### Go (arrays)
```go
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// primes is of type [6]int
	primes := [...]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

```
**Output**:
```plaintext
Hello World
[Hello World]
[2 3 5 7 11 13]
```

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

## Slice length and capacity
Slices are defined by a pointer to the start of their data and two other fields: **Length** and **Capacity**. The length of a slice gives users knowledge of the accesible/useful data in a slice. The capacity is a property used primarily by the garbage collector to determine the available space in the slice before needing to allocate a new slice when adding elements to the end of a slice with `append` (more on that later). 

Use the builtins `len` and `cap` to get the length and capacity of a slice, respectively:
```go
var a [10]int
len(a[:])  // 10
len(a[:0]) // 0
len(a[:8]) // 8
len(a[9:]) // 1

cap(a[:])  // 10
cap(a[:0]) // 10
cap(a[:8]) // 10
cap(a[9:]) // 1
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
Slices can contain any type, including other slices... and those slices can contain more slices, and so on:
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


### Python (slices)
```python
s = [2, 3, 5, 7, 11, 13]

s = s[1:4]
print(s)

s = s[:2]
print(s)

s = s[1:]
print(s)


```
### Go (slices)
```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

```
**Output**:
```plaintext
[3 5 7]
[3 5]
[5]
```

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
### Python (append)
```python
s = []
print(s)

s.append(0)
print(s)

s.append(1)
print(s)

# Python does not allow adding more than one item at a time.
s.extend([2, 3, 4])
print(s)

s.extend([5, 6, 7])
print(s)
```
### Go (append)
```go
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s)

	// append works on nil slices.
	s = append(s, 0)
	fmt.Println(s)

	// The slice grows as needed.
	s = append(s, 1)
	fmt.Println(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	fmt.Println(s)

	// We can also append a list to a list.
	g := []int{5, 6, 7}
	s = append(s, g...)
	fmt.Println(s)
}

```
**Output**:
```plaintext
[]
[0]
[0 1]
[0 1 2 3 4]
[0 1 2 3 4 5 6 7]
```

# Range
The `range` form of the `for` loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

You can skip the index or value by assigning to _.
```go
for i, _ := range pow
for _, value := range pow

// If you only want the index, you can omit the second variable.
for i := range pow
```
### Python (range)
```python

pow = [1, 2, 4, 8, 16, 32, 64, 128]

for i, v in enumerate(pow):
    print(f"2**{i} = {v}")

for i in range(len(pow)):
    print(f"2**{i} = {pow[i]}")

for v in pow:
    print(f"pow {v}")

```
### Go (range)
```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	for i := range pow {
		fmt.Printf("2**%d = %d\n", i, pow[i])
	}
	for _, v := range pow {
		fmt.Printf("pow %d\n", v)
	}
}

```
**Output**:
```plaintext
2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
2**4 = 16
2**5 = 32
2**6 = 64
2**7 = 128
2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
2**4 = 16
2**5 = 32
2**6 = 64
2**7 = 128
pow 1
pow 2
pow 4
pow 8
pow 16
pow 32
pow 64
pow 128
```

# Maps
A map maps keys to values.

The zero value of a map is `nil`. A `nil` map has no keys, nor can keys be added.

The `make` function returns a map of the given type, initialized and ready for use.

## Mutating Maps
Insert or update an element in map m:
```go
m[key] = elem
```

Retrieve an element:
```go
elem = m[key]
```

Delete an element:
```go
delete(m, key)
```

Test that a key is present with a two-value assignment:
```go
elem, ok = m[key]
```

If key is in m, ok is true. If not, ok is false.

If key **is not** in the map, then elem is the **zero value** for the map's element type.

Note: If elem or ok have not yet been declared you could use a short declaration form:
```go
elem, ok := m[key]
```
### Python (maps)
```python
ages = {
    "Sarah":        32,
    "Billy":        12,
    "Jeremiah":     99,
    "John Baptist": 47,
}
print(ages["Sarah"])
billyAge, billyPresent = 0, False
if "Billy" in ages:
    billyAge, billyPresent = ages["Billy"], True

print(billyAge, billyPresent)

x13Age, x13Present = 0, False
if "x13" in ages:
    x13Age, x13Present = ages["x13"], True

print(x13Age, x13Present)

ages["Faustus"] = 66
print(ages)
```
### Go (maps)
```go
package main

import "fmt"

func main() {
	ages := map[string]int{
		"Sarah":        32,
		"Billy":        12,
		"Jeremiah":     99,
		"John Baptist": 47,
	}
	fmt.Println(ages["Sarah"])

	billyAge, billyPresent := ages["Billy"]
	fmt.Println(billyAge, billyPresent)

	x13Age, x13Present := ages["x13"]
	fmt.Println(x13Age, x13Present)

	ages["Faustus"] = 66
	fmt.Println(ages)
}

```
**Output**:
```plaintext
32
12 true
0 false
map[Billy:12 Faustus:66 Jeremiah:99 John Baptist:47 Sarah:32]
```

# Pointers
TODO
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

### Python (inline-functions)
```python
def SuperRandom(normalRandom:lambda:int) -> int:
    """SuperRandom returns a super random number using a
    not so random source function "normalRandom".
    (this is not actually more random!)"""
    superrand = 12345678
    for _ in range(3):
        rand = normalRandom()
        superrand = superrand*7 + rand*31
    return superrand

# Python's inline function creation uses
# the lambda keyword, though its limited to
# one-liners so we can't really use it here.
# We do with a simple def.
a = 1
def notSoRand()->int:
    global a
    a = a * 7
    return a

print(notSoRand(), notSoRand())

superrand1 = SuperRandom(notSoRand)
superrand2 = SuperRandom(notSoRand)
superrand3 = SuperRandom(notSoRand)

print(superrand1, superrand2, superrand3)





```
### Go (inline-functions)
```go
package main

import "fmt"

func main() {
	a := 1
	notSoRand := func() int {
		a = a * 7
		return a
	}
	call1 := notSoRand()
	call2 := notSoRand()
	fmt.Println("calling function yields different results:", call1, call2)

	superrand1 := SuperRandom(notSoRand)
	superrand2 := SuperRandom(staticRandom)
	fmt.Println("a function can take another function as argument:", superrand1, superrand2)
}

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

func staticRandom() int {
	return 4
}

```
**Output**:
```plaintext
calling function yields different results: 7 49
a function can take another function as argument: 4236130605 4234574622
```

