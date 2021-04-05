package main

func SumAll(numbersToSum ...[]int) []int {

	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, number := range numbersToSum {
		sums[i] = Sum(number)
	}
	// fmt.Printf(sums)
	return sums
}
