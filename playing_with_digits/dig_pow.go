package playing_with_digits

import "math"

func DigPow(n, p int) int {
	if n == 0 {
		return -1
	}

	remainder := float64(n)
	numDigits := 1 + int(math.Log10(remainder))

	sum := 0.0
	for i := 0; i < numDigits; i++ {
		currentDigitPosition := float64(numDigits - 1 - i)
		currentDigitPow10 := math.Pow(10, currentDigitPosition)

		currentDigit := math.Floor(remainder / currentDigitPow10)
		sum += math.Pow(currentDigit, float64(p+i))

		remainder -= currentDigit * currentDigitPow10
	}

	k := sum / float64(n)
	if math.Trunc(k) == k {
		return int(k)
	}

	return -1
}
