package uuid

import (
	"strconv"
	"strings"
	"time"
)

// var autoIncrement int64 = 1 // 设置当前机器的自增ID

func GetSnowflakeId(autoIncrement int64) int64 {
	// 获取41位时间戳二进制字符串，2023-01-01作为开始时间
	startTime, _ := time.Parse("2006-01-02", "2023-01-01")
	diffTime := time.Now().Sub(startTime).Milliseconds() // 获取当前时间和2023-01-01之间的毫秒时间戳
	diffTimeStr := handleNum(diffTime, 41)

	// 假设机器编号是2
	serialNumber := int64(2)
	serialNumberStr := handleNum(serialNumber, 10)

	autoIncrementStr := handleNum(autoIncrement, 12)
	total := "0" + diffTimeStr + serialNumberStr + autoIncrementStr
	uid, _ := strconv.ParseInt(total, 2, 64) // 二进制字符串转位10进制的64位整数

	// autoIncrement++
	return uid
}

// 将int64转换成指定长度的二进制字符串
func handleNum(num int64, targetlength int) string {
	str := strconv.FormatInt(num, 2)
	strlen := len(str)
	if strlen > targetlength {
		// error
		return ""
	} else if strlen == targetlength {
		return str
	} else {
		diffL := targetlength - strlen
		x := make([]string, diffL)
		for i := range x {
			x[i] = "0"
		}
		xstr := strings.Join(x, "")
		return xstr + str
	}
}
