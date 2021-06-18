package decode_morse_code_2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/wonderbird/katas/decode_morse_code_2"
)

var _ = Describe("DecodeMorse", func() {
	PIt("Codewars: Example from description", func() {
		Expect(DecodeMorse(DecodeBits("1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011"))).To(Equal("HEY JUDE"))
	})
})
