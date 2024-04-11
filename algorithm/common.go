package algorithm

import (
	"errors"
	"math"
)

// 获取数据数组中的最大值
func GetInt64ArrayMaxValue(numbers []int64) int64 {
	if len(numbers) == 0 {
		return -1
	}
	var max int64 = numbers[0]
	for i := 0; i < len(numbers); i++ {
		if max < numbers[i] {
			max = numbers[i]
		}
	}
	return max
}

// 获取数据数组中的最大值和最小值
func GetInt64ArrayMaxAndMinValue(numbers []int64) (min, max int64) {
	if len(numbers) == 0 {
		return -1, -1
	}
	min, max = numbers[0], numbers[0]
	for i := 0; i < len(numbers); i++ {
		if max < numbers[i] {
			max = numbers[i]
		}
		if min > numbers[i] {
			min = numbers[i]
		}
	}

	return
}

// 判断整数位数
/*
math.Abs用于获取数字的绝对值，以防传入的数字是负数。
math.Log10计算了数字的对数（以10为底）。
math.Floor返回当前数据的最大整数。
*/
func CountDigits(n int) int {
	return int(math.Floor(math.Log10(float64(math.Abs(float64(n)))))) + 1
}

// 获取 整数数据number上位数digit（个、十、百、千、万等）的数据
/*
 digit = 1，表示取个位上的数据
 digit = 10，表示取十位上的数据
*/
func GetDigitValue(number, digit int64) (uint8, error) {
	// 判断digit是否是10的整数倍，如果不是，报错
	if digit == 0 || !IsInteger(math.Log10(float64(digit))) {
		return 0, errors.New("digit error")
	}

	// 比如传入的number=178，dight=10，我们要去十位上的数字7
	md := digit * 10
	m := number % md // 取从digit位开始的数字，178 % 100 = 78
	v := m / digit   // 78 / 10 = 7
	return uint8(v), nil
}

// 判断浮点数是否位整数
/*
math.Trunc 返回浮点数的整数部分。
number == math.Trunc(number)  ：判断number是否等于其整数部分
math.IsInf 判断数字是否无穷大
*/
func IsInteger(number float64) bool {
	return math.IsInf(number, 0) || number == math.Trunc(number)
}
