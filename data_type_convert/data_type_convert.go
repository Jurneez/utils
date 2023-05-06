package data_type_convert

import (
	"bytes"
	"encoding/json"
	"fmt"
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

// 将json字符串进行格式化
/*
比如：传入{"Name":"amy","Age":18}
输出：
{
	"Name":"amy",
	"Age":18
}
*/
func JsonString(str string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(str), "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", str)
	}

	return out.String()
}
