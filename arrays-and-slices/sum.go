package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToTails ...[]int) []int {
	var sumTails []int

	for _, numbers := range numbersToTails {
		if len(numbers) == 0 {
			sumTails = append(sumTails, 0)
			continue
		}
		tail := numbers[1:]
		sumTails = append(sumTails, Sum(tail))
	}
	return sumTails
}
