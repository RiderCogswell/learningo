package main 

func Sum(numbers []int) int { // arrays lengths must be declared on init, so instead we use a slice for flexibility
	sum := 0
	// range iterates over an array and returns the index, and the value
	for _, number := range numbers { // we are choosing to ignore the iterator by using the blank identifier => _
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNums := len(numbersToSum)
	sums := make([]int, lengthOfNums)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}