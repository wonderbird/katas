package highest_rank_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/wonderbird/katas/highest_rank"
)

var _ = Describe("Tests", func() {
	Describe("TDD tests", func() {
		TestHighestRank := func(input []int, expected int) {
			Expect(highest_rank.HighestRank(input)).To(Equal(expected))
		}

		DescribeTable("Single element in array",
			TestHighestRank,
			Entry("7", []int{7}, 7),
			Entry("2", []int{2}, 2),
		)

		DescribeTable("Two ascending elements in array",
			TestHighestRank,
			Entry("1, 2", []int{1, 2}, 2),
		)
	})

	PDescribe("Codewars tests", func() {
		It("Codewars test 1: 12, 10, 8, 12, 7, 6, 4, 10, 12", func() {
			Expect(highest_rank.HighestRank([]int{12, 10, 8, 12, 7, 6, 4, 10, 12})).To(Equal(12))
		})
	})
})
