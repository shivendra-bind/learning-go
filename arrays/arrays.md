Arrays allow you to store multiple elements of the same type in a variable in a particular order

Arrays have a fixed capacity which you define when you declare the variable. We can initialize an array in two ways:

- [N]type{value1, value2, ..., valueN} e.g. `numbers := [5]int{1, 2, 3, 4, 5}`
- [...]type{value1, value2, ..., valueN} e.g. `numbers := [...]int{1, 2, 3, 4, 5}`

- *%v* placeholder to print the "default" format, which works well for arrays

### range
`range` iterates over an array. On each iteration, range returns two values - the index and the value

### Creating a slice with make
Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array:

```go
a := make([]int, 5)  // len(a)=5
To specify a capacity, pass a third argument to make:

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

### append

`append` function which takes a slice and a new value, then returns a new slice with all the items in it

### slice

`slice[low:high]` syntax for slicing an array
- it is inclusive of index