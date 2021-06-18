package decode_morse_code_2

import (
	"github.com/wonderbird/katas/decode_morse_code_1"
)

// You may get original char by morse code like this: MORSE_CODE[char]
type CodeLengthPair struct {
	Code   rune
	Length int
}

func DecodeBits(bits string) string {
	currentCode := '0'
	currentLength := 0
	codeLengths := []CodeLengthPair{}

	for _, bit := range bits {
		if bit == currentCode {
			currentLength++
		} else {
			codeLengths = append(codeLengths, CodeLengthPair{
				currentCode,
				currentLength,
			})
			currentCode = bit
			currentLength = 1
		}
	}
	codeLengths = append(codeLengths, CodeLengthPair{
		currentCode,
		currentLength,
	})

	if codeLengths[0].Code == '0' {
		codeLengths = codeLengths[1:]
	}
	timeUnit := codeLengths[0].Length

	morseCode := ""
	for _, codeLength := range codeLengths {
		if codeLength.Code == '1' {
			if codeLength.Length/timeUnit == 1 {
				morseCode += "."
			} else {
				morseCode += "-"
			}
		}
	}

	return morseCode
}

func DecodeMorse(morseCode string) string {
	return decode_morse_code_1.DecodeMorse(morseCode)
}
