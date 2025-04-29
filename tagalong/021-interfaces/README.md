# Interfaces
Interfaces in Go provide a way to define behaviour for a undefined type that has a set of methods. This is how the ubiquitous [`io.Reader`](https://pkg.go.dev/io#Reader) interface is implemented in the Go standard library.

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

`io.Reader` is world famous. It appears in nearly every single Go project, either directly or indirectly. This is because it is implemented or used by a large part of the standard library to represent any of the following:

- [A file](https://pkg.go.dev/os#File.Read)
- [A byte buffer](https://pkg.go.dev/bytes#Buffer.Read)
- [A TCP network stream](https://pkg.go.dev/net#TCPConn.Read)
- [An HTTP body](https://pkg.go.dev/net/http#Request)
- [An HTTP Multipart form file](https://pkg.go.dev/mime/multipart#Part.Read)
- [A Zip file header](https://pkg.go.dev/archive/zip#File.OpenRaw)
- [A Targz entry](https://pkg.go.dev/archive/tar#Reader.Read)
- And so much more

To get an idea of how powerful this is, imagine you wrote an algorithm to parse a special format- maybe it's a new programming language you are writing. You only need to define one function that takes in an `io.Reader` and write the logic for it once. From then on that function can take in any of the aforementioned data streams (a OS file, HTTP Body, a zipped archive, a raw TCP stream, etc.)

```go
func ParseFormat(r io.Reader) Format
```

The underlying function needs absolutely no knowledge about the underlying implementation. It relieves the programmer of a lot of thinking. 

The effect is also multiplicative- since the type is so ubiquitous in the standard library most open source libraries also implement the interface for their types leading to practically infinite combinations of ways to call functions that receive an `io.Reader` type.

## Comparison with other languages
Interfaces are the way Go goes about structural typing polymorphism. Similar concepts exist in most common languages but usually by other names:
- Python's `typing.Protocol`
- Typescript's `protocol`

Other languages do something similar to Go but require additional work of specifying the interface that is implemented. This is known as Nominal typing polymorphism

- Java's `interface`
- Rust's `trait`
- C#'s `interface`
- Python's `typing.ABC`

## Observations

- **Keep it simple:** Strive for simplicity to see your interfaces suceed in use. The Go standard library's most succesful interfaces have a single method, these are the famous "single method interfaces". Even the [embedded system interfaces of TinyGo](https://github.com/tinygo-org/drivers/blob/release/i2c.go#L5) are single method interfaces!
- **Automatic satisfaction**: In more ways than one. You don't need to write `implements` declarations explcitly saying what interface is being implemented by a method. In Go interfaces are implemented as soon as you write out the last method required to implement a method set.
- **Type assertion:** In Go you can convert an interface to its concret type. Always use the two parameter return to avoid panics in conversion: `c, ok := v.(MyInterface)`
- **Empty interface or any:** The any interface is implemented by all types. This type is useless unless using `reflect` to introspect it or trying type assertions, like `fmt.Println` does.
- **Nil panics:** Calling a method on a nil interface will panic.
- **Nil interface vs nil subtype:** A very confusing pitfall. An interface is nil as long as it has not subtype assigned. As soon as there's a subtype associated to it a nil check will return false even though the underlying pointer is nil.

```go
var a any
var i *int
fmt.Println(a, i, a==nil) // <nil> <nil> true
a = i
fmt.Println(a, a==nil) // <nil> false
```