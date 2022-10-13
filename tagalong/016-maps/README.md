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