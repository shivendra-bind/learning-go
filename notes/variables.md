- `static` typed language
- requires data `type` at variable declaration
- can infer `type` from *initialized variable*

## keywords
1. `var` declares one or more variables
2. `const` declares a read-only variable (cannot be reassigned)

### example
```go
var a = "initial"
var b, c int = 1, 2
var d = true
var e int // default value is 0
f := "apple" // shorthand for var f string = "apple"
```

```go
const Pi = 3.14
// numeric constant has no type until it’s given one, such as by an explicit cast.
```

## Data Types
- `Numeric`
    - `int`
        - `int8`
        - `int16`
        - `int32`
        - `int64`
    - `float`
        - `float32`
        - `float64`
    - `complex`
        - `complex64`
        - `complex128`
- `String`
- `Boolean`

- `Derived`
    - `Pointer`
    - `Array`
    - `Structure`
    - `Map`
    - `Interface`


1. In `Go` if a symbol (variables, types, functions et al) starts with a lowercase symbol then it is private outside the package it's defined in.    
2. In `Go`, when you call a `function` or a `method` the arguments are _**copied**_.
3.  Pointers let us *point* to some values and then let us change them