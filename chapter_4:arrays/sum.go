package main

func Sum(values []int) int {
	sum := 0

	for _, number := range values {
		sum += number
	}
	return sum
}
