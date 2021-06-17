package playing_with_digits

import "math"

func DigPow(n, p int) int {
	// TODO can we leave this check out?
	if n == 0 {
		return -1
	}

	nf := float64(n)

	sum := 0.0
	numDigits := 1 + int(math.Log10(nf))
	remainder := nf
	for i := 0; i < numDigits; i++ {
		currentDigitPosition := float64(numDigits - 1 - i)
		currentDigitPow10 := math.Pow(10, currentDigitPosition)

		currentDigitValue := math.Floor(remainder / currentDigitPow10)
		sum += math.Pow(currentDigitValue, float64(p+i))

		remainder -= currentDigitValue * currentDigitPow10
	}

	k := sum / float64(n)
	if math.Trunc(k) == k {
		return int(k)
	}

	return -1
}
