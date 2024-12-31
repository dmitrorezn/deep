
Deep Package

The deep package provides utility functions to perform deep equality checks on slices and maps in Go. It leverages the reflect package and custom logic to handle nested structures efficiently.

Features
```
Deep Slice Comparison: Compare slices for equality, including nested slices.
Deep Map Comparison: Compare maps for equality, including nested maps.
Type Flexibility: Supports generic slice and map types for comparison.
```

Functions
EqualSlices
Compares two slices for deep equality.

Signature:
```go
func EqualSlices[S interface{ ~[]E }, E comparable](s1 S, s2 S) bool
```
Usage:

```go
s1 := []int{1, 2, 3}
s2 := []int{1, 2, 3}
equal := deep.EqualSlices(s1, s2) // true
```
EqualMaps
Compares two maps for deep equality.

Signature:

```go
func EqualMaps[M1 interface{ ~map[K]V }, M2 interface{ ~map[K]V }, K comparable, V comparable](m1 M1, m2 M2) bool
```
Usage:

```go
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"a": 1, "b": 2}
equal := deep.EqualMaps(m1, m2) // true
```
equalValues
A helper function for internal use that compares two values using reflection.

Signature:

```go
func equalValues[V comparable](v1, v2 V) bool
```
This function checks:

Nested slices or arrays for equality.
Nested maps for equality.
Basic value comparison for other types.
Example
Here's a complete example demonstrating how to use the deep package:

```go
package main

import (
    "fmt"
    "github.com/dmitrorezn/deep"
)

func main() {
    // Compare slices
    s1 := []any{1, map[int]int{1:1}, []int{1, 2}}
    s2 := []any{1, map[int]int{1:1}, []int{1, 2}}
    fmt.Println("Slices Equal:", deep.EqualSlices(s1, s2)) // Output: true

	// Compare maps
	m1 := map[string]any{"a": 1, "b": 2}
	m2 := map[string]any{"a": 1, "b": 2}
	fmt.Println("Maps Equal:", deep.EqualMaps(m1, m2)) // Output: true
}
```