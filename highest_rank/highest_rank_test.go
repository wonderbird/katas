package highest_rank_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/wonderbird/katas-in-go/highest_rank"
	"testing"
)

var _ = Describe("Tests", func() {
	Describe("TDD tests", func() {
		DescribeTable("Single element in array",
			RunTestForBothSolutions,
			CreateEntryWithAutogeneratedDescription([]int{7}, 7),
			CreateEntryWithAutogeneratedDescription([]int{2}, 2),
		)

		DescribeTable("Only ascending elements in array with single occurrence",
			RunTestForBothSolutions,
			CreateEntryWithAutogeneratedDescription([]int{1, 2}, 2),
			CreateEntryWithAutogeneratedDescription([]int{-1, 1}, 1),
			CreateEntryWithAutogeneratedDescription([]int{1, 2, 6, 9}, 9),
		)

		DescribeTable("Mixed sequence elements in array with single occurrence",
			RunTestForBothSolutions,
			CreateEntryWithAutogeneratedDescription([]int{32, 16}, 32),
			CreateEntryWithAutogeneratedDescription([]int{32, 56, 12, 65, 16}, 65),
		)

		DescribeTable("One element repeated twice all others once",
			RunTestForBothSolutions,
			CreateEntryWithAutogeneratedDescription([]int{2, 2}, 2),
			CreateEntryWithAutogeneratedDescription([]int{2, 2, 3}, 2),
			CreateEntryWithAutogeneratedDescription([]int{0, -2, -2, 3}, -2),
			CreateEntryWithAutogeneratedDescription([]int{0, 3, 3}, 3),
		)

		DescribeTable("Special cases",
			RunTestForBothSolutions,
			CreateEntryWithAutogeneratedDescription([]int{0, 0, 3, 3}, 3),
			CreateEntryWithAutogeneratedDescription([]int{-7, -7, -7, 6, 6}, -7),
			CreateEntryWithAutogeneratedDescription([]int{-5, 1, 13, -7, -5, 8, -5, 13}, -5),
		)

		DescribeTable("Examples from kata description",
			RunTestForBothSolutions,
			CreateEntryWithAutogeneratedDescription([]int{12, 10, 8, 12, 7, 6, 4, 10, 12}, 12),
			CreateEntryWithAutogeneratedDescription([]int{12, 10, 8, 12, 7, 6, 4, 10, 12, 10}, 12),
			CreateEntryWithAutogeneratedDescription([]int{12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10}, 3),
		)
	})

	Describe("Codewars tests", func() {
		It("Codewars test 1: 12, 10, 8, 12, 7, 6, 4, 10, 12", func() {
			RunTestForBothSolutions([]int{12, 10, 8, 12, 7, 6, 4, 10, 12}, 12)
		})
	})
})

var benchmarkInput = []int{4, 1, 8, 46, 7, 0, 13, 47, 1, 13, 23, 91, 13, -56, 12, 0}

func BenchmarkHighestRank(b *testing.B) {
	for i := 0; i < b.N; i++ {
		highest_rank.HighestRank(benchmarkInput)
	}
}

func BenchmarkAlternativeSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		highest_rank.AlternativeSolution(benchmarkInput)
	}
}

func RunTestForBothSolutions(input []int, expected int) {
	RunTestForHighestRank(input, expected)
	RunTestForAlternativeSolution(input, expected)
}

func RunTestForHighestRank(input []int, expected int) {
	Expect(highest_rank.HighestRank(input)).To(Equal(expected))
}

func RunTestForAlternativeSolution(input []int, expected int) {
	Expect(highest_rank.AlternativeSolution(input)).To(Equal(expected))
}

func CreateEntryWithAutogeneratedDescription(input []int, expected int) TableEntry {
	return Entry(fmt.Sprintf("%v -> %d", input, expected), input, expected)
}
