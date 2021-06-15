package highest_rank

import "sort"

func HighestRank(nums []int) int {
	sort.Ints(nums)

	mostFrequentNumber := 0
	highestFrequency := 0
	currentNumber := 0
	currentFrequency := 0

	for _, num := range nums {
		if num != currentNumber {
			currentNumber = num
			currentFrequency = 1
		} else {
			currentFrequency++
		}

		if currentFrequency >= highestFrequency {
			mostFrequentNumber = num
			highestFrequency = currentFrequency
		}
	}

	return mostFrequentNumber
}
