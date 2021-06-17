package playing_with_digits_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"highest_rank/playing_with_digits"
)

func dotest(n, p int, exp int) {
	var ans = playing_with_digits.DigPow(n, p)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(89, 1, 1)
		dotest(92, 1, -1)
	})
})
