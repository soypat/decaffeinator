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