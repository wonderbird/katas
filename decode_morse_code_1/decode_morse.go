package decode_morse_code_1

import "strings"

func DecodeMorse(morseCode string) string {
	morseCode = strings.TrimSpace(morseCode)

	encodedWords := strings.Split(morseCode, "   ")
	var decodedWords []string

	for _, encodedWord := range encodedWords {
		decodedWords = append(decodedWords, DecodeWord(encodedWord))
	}
	sentence := strings.Join(decodedWords, " ")

	return sentence
}

func DecodeWord(morseCode string) string {
	decoded := ""

	splitCodes := strings.Split(morseCode, " ")
	for _, code := range splitCodes {
		decoded += MORSE_CODE[code]
	}

	return decoded
}

var MORSE_CODE = map[string]string{
	".-":        "A",
	"-...":      "B",
	"-.-.":      "C",
	"-..":       "D",
	".":         "E",
	"..-.":      "F",
	"--.":       "G",
	"....":      "H",
	"..":        "I",
	".---":      "J",
	"-.-":       "K",
	".-..":      "L",
	"--":        "M",
	"-.":        "N",
	"---":       "O",
	".--.":      "P",
	"--.-":      "Q",
	".-.":       "R",
	"...":       "S",
	"-":         "T",
	"..-":       "U",
	"...-":      "V",
	".--":       "W",
	"-..-":      "X",
	"-.--":      "Y",
	"--..":      "Z",
	".-.-.-":    ".",
	"-.-.--":    "!",
	"...---...": "SOS",
}
