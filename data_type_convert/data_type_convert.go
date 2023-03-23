package data_type_convert

import (
	"strconv"
)

// string to int
func StringToInt(str string) int {
	int, _ := strconv.Atoi(str)
	return int
}

// 输出含有特殊字符的字符串
func PrintSpeicalString(str string) string {
	return strconv.Quote(str)
}
