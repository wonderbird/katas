package decode_morse_code_1_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	. "github.com/wonderbird/katas-in-go/decode_morse_code_1"
)

var _ = Describe("Test DecodeMorse", func() {
	Describe("WHEN morse code is empty", func() {
		It("THEN return empty string", func() {
			Expect(DecodeMorse("")).To(Equal(""))
		})
	})

	DescribeTable("WHEN morse code is a single letter",
		AssertInputIsDecodedCorrectly,
		CreateEntryWithAutogeneratedDescription(".", "E"),
		CreateEntryWithAutogeneratedDescription(".-", "A"),
		CreateEntryWithAutogeneratedDescription("---", "O"),
	)

	DescribeTable("WHEN morse code is a single word",
		AssertInputIsDecodedCorrectly,
		CreateEntryWithAutogeneratedDescription(". --. --.", "EGG"),
		CreateEntryWithAutogeneratedDescription("... --- ...", "SOS"),
	)

	DescribeTable("WHEN morse code contains several words",
		AssertInputIsDecodedCorrectly,
		CreateEntryWithAutogeneratedDescription("-... .-. .. -. --.   . --. --.", "BRING EGG"),
	)

	It("Codewars Example from description", func() {
		Expect(DecodeMorse(".... . -.--   .--- ..- -.. .")).To(Equal("HEY JUDE"))
	})
})

func AssertInputIsDecodedCorrectly(input, expected string) {
	Expect(DecodeMorse(input)).To(Equal(expected))
}

func CreateEntryWithAutogeneratedDescription(input, expected string) TableEntry {
	return Entry(fmt.Sprintf("%q -> %q", input, expected), input, expected)
}
