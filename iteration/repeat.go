package iteration

func Repeat(char string, iteration int) string {
	var result string
	for i := 0; i < iteration; i++ {
		result = result + char
	}
	return result
}
