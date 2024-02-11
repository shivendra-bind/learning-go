package arrays

func Sum(arr []int) int {
	sum := 0
	for _, element := range arr {
		sum += element
	}
	return sum
}
