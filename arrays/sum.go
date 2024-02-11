package arrays

func Sum(arr []int) int {
	sum := 0
	for _, element := range arr {
		sum += element
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lenArgs := len(numbersToSum)
	sums := make([]int, lenArgs)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
