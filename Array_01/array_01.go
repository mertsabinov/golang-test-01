package array_01

func Sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func SumAll(numberList ...[]int) []int {
	// Write a function that returns a number list
	// Return the sum of the slices entered
	// add to a new list
	var sums []int
	for _, n := range numberList {
		sums = append(sums, Sum(n))
	}
	return sums
}
