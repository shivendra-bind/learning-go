# Testing in go

### rules
1. File name format should be like `filename_test.go`
2. Test function must start with with word `Test`
3. Test function takes arg `t *testing.T`

### Table driven testing

Table driven tests are useful when you want to build a list of test cases that can be tested in the same manner.
```go
func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}
```