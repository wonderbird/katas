package playing_with_digits_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"highest_rank/playing_with_digits"
)

func dotest(n, p int, exp int) {
	var ans = playing_with_digits.DigPow(n, p)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test DigPow", func() {

	DescribeTable("Inconclusive cases using n = 0",
		dotest,
		FEntry("0, 0 => -1 (undefined: 0^0)", 0, 0, -1),
		Entry("0, 1 => -1 (inconclusive: 0^1 = 0 = 1 * 0 = ... * 0)", 0, 1, -1),
		Entry("0, 5 => -1 (inconclusive: 0^5 = 0 = 1 * 0 = ... * 0)", 0, 5, -1),
	)

	DescribeTable("Cases using p = 0",
		dotest,
		Entry("1, 0 => 1  (1^0 = 1  = 1*1)", 1, 0, 1),
		Entry("3, 0 => -1 (3^0 = 1 != k*3)", 3, 0, -1),
	)

	DescribeTable("Single digit cases using p = 1",
		dotest,
		Entry("1, 1 => 1  (1^1 = 1 = 1*1)", 1, 1, 1),
		Entry("5, 1 => 1  (5^1 = 5 = 1*5)", 5, 1, 1),
	)

	DescribeTable("Double digit cases using p = 1",
		dotest,
		Entry("10, 1 => -1  (1^1 + 0^2 = 1 < 1*10)", 10, 1, -1),
		Entry("11, 1 => -1  (1^1 + 1^2 = 2 < 1*11)", 11, 1, -1),
		Entry("89, 1 => 1  (8^1 + 9^2 = 89 < 1*89)", 89, 1, 1),
	)

	PIt("(codewars) should handle basic cases", func() {
		dotest(92, 1, -1)
		dotest(695, 2, 2)
		dotest(46288, 3, 51)
	})
})
