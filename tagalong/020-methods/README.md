# Methods
If coming from a language with rich OOP features Go may begin to feel sparse at this point. In Go what we call methods don't really bring any added functionality to the table over functions. A method is just that, a function with an extra "receiver" argument.

We declare methods by writing `func` followed by parentheses containing the receiver identifier and type, the rest of the function declaration follows as normal after that point:

```go
func (self RectangleClass) Area() int
```

Some observations:

- **Package level scope:** Methods are declared at the same scope as package level functions.

- **Sub-package namespacing:** Methods can only be called on the type using dot notation. Thus they are useful for not cluttering the package namespace.

- **Exporting: public/private**: Just like with all package-level identifiers, the method is exported if it starts with a upper case letter

- **Pointer receivers:** We can define our receiver to be a pointer type or non-pointer type

- **Receiver type limits:** You can only define methods on type defined in the local package. If you so desire to have your own methods next to the foreign type you may decide to go the struct embedding route (see Go's struct embedding).

## Interfaces
Methods are revealed to be an unexpecedly powerful feature when interfaces are introduced. Interfaces are the equivalent of typing.Protocol in Python.

