package time

import (
	"testing"
	"tool/utils/common"
)

func Test_TimeConvTimestemp(t *testing.T) {
	t.Log(TimeConvTimestemp(common.Time_Format_Second, "2021-10-01 00:00:00"))
}
