package algorithm

import (
	"math"
	"testing"
)

func Test_CountDigits(t *testing.T) {
	t.Log(math.Abs(-1092)) // 1092
	t.Log(math.Log10(10))  // 3.0382226383687185
}

func Test_GetDigitValue(t *testing.T) {
	digit, _ := GetDigitValue(187, 1)
	t.Log(digit)
}
