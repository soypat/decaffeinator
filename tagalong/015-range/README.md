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