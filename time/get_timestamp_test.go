package time

import (
	"testing"
	"time"
)

func Test_GetMonthStartTimeStamp(t *testing.T) {
	t.Logf("月初时间戳： %d\n", GetMonthStartTimeStamp(2021, time.Month(11)))

	t.Logf("月末时间戳： %d\n", GetMonthEndTimeStamp(2021, time.Month(11)))
}
