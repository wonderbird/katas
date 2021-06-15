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

// AlternativeSolution was implemented after the solution above has
// been completed using TDD.
//
// I was expecting that this solution is more efficient, because
// it uses only two loops of size N (O(2N)).
//
// The benchmark tests show that the alternative solution is half
// as fast as the original solution, which I expect to be
// O(N log(N) + N).
//
// I guess that the map introduces another O(N log(N)) which is
// nested into the loop. So the efficiency of the alternative
// solution is likely to be in the range of O(N * N log(N))
// for large N and O(N log(N)) for small N.
//
// Benchmark results on my machine:
//   goos: darwin
//   goarch: amd64
//   pkg: github.com/wonderbird/katas/highest_rank
//   cpu: Intel(R) Core(TM) i7-6920HQ CPU @ 2.90GHz
//   BenchmarkHighestRank
//   BenchmarkHighestRank-8           	 5364417	       221.8 ns/op
//   BenchmarkAlternativeSolution
//   BenchmarkAlternativeSolution-8   	 1326914	       900.7 ns/op
//   PASS
//   ok  	github.com/wonderbird/katas/highest_rank	3.682s
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
