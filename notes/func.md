# func

keyword to declare function

all args have type

```go
func fn(a int, b string, c string) {}
```

if function returns anything then return type should be declaired

```go
func fn(a, b, c string) int {
    return len(a + b + c) 
}
```

named return

```go
func fn(a, b, c string) (joined string) {
    joined = a + b + c
    return
}
```
