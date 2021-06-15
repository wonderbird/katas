package highest_rank

import "sort"

func AlternativeSolution(nums []int) int {
	frequencyByNumber := map[int]int{}

	for _, num := range nums {
		frequencyByNumber[num]++
	}

	highestFrequency := 0
	mostFrequentNumber := 0
	for num, freq := range frequencyByNumber {
		if (freq > highestFrequency) || (freq == highestFrequency && num > mostFrequentNumber) {
			highestFrequency = freq
			mostFrequentNumber = num
		}
	}

	return mostFrequentNumber
}

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
