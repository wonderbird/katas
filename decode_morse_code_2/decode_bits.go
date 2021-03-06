package decode_morse_code_2

import (
	"github.com/wonderbird/katas-in-go/decode_morse_code_1"
	"math"
	"strings"
)

// You may get original char by morse code like this: MORSE_CODE[char]
type CodeLengthPair struct {
	Code   rune
	Length int
}

func DecodeBits(bits string) string {
	trimmedBits := strings.Trim(bits, "0")

	codeLengths := convertBitsToCodeLengthPairs(trimmedBits)
	timeUnit := findLengthOfShortestZeroBurst(codeLengths)
	morseCode := convertCodeLengthPairsToMorseCode(codeLengths, timeUnit)

	return morseCode
}

func convertBitsToCodeLengthPairs(bits string) (codeLengths []CodeLengthPair) {
	currentCode := '1'
	currentLength := 0

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

	return
}

func findLengthOfShortestZeroBurst(codeLengths []CodeLengthPair) int {
	timeUnit := codeLengths[0].Length
	for _, codeLength := range codeLengths {
		timeUnit = int(math.Min(float64(timeUnit), float64(codeLength.Length)))
	}
	return timeUnit
}

func convertCodeLengthPairsToMorseCode(codeLengths []CodeLengthPair, timeUnit int) string {
	morseCode := ""
	for _, codeLength := range codeLengths {
		if codeLength.Code == '1' {
			if codeLength.Length/timeUnit == 1 {
				morseCode += "."
			} else {
				morseCode += "-"
			}
		} else if codeLength.Code == '0' {
			if codeLength.Length/timeUnit == 3 {
				morseCode += " "
			} else if codeLength.Length/timeUnit == 7 {
				morseCode += "   "
			}
		}
	}
	return morseCode
}

func DecodeMorse(morseCode string) string {
	return decode_morse_code_1.DecodeMorse(morseCode)
}
