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

func SumAllAppend(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			continue
		}
		sums = append(sums, Sum(numbers[1:]))
	}

	return sums
}
