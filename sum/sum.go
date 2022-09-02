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
	var sums []int // init a slice

	for _, numbers := range numbersToSum { // loop through
		sums = append(sums, Sum(numbers)) // append each sum, 
	}

	return sums
}