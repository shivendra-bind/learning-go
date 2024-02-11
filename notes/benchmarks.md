# Benchmarks 
Functions of the form

```go
func BenchmarkXxx(*testing.B)
```
are considered benchmarks, and are executed by the "go test" command when its -bench flag is provided. Benchmarks are run sequentially