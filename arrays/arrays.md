Arrays allow you to store multiple elements of the same type in a variable in a particular order

Arrays have a fixed capacity which you define when you declare the variable. We can initialize an array in two ways:

- [N]type{value1, value2, ..., valueN} e.g. `numbers := [5]int{1, 2, 3, 4, 5}`
- [...]type{value1, value2, ..., valueN} e.g. `numbers := [...]int{1, 2, 3, 4, 5}`

- *%v* placeholder to print the "default" format, which works well for arrays

### range
`range` iterates over an array. On each iteration, range returns two values - the index and the value