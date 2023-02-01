package slice

import (
	"testing"
)

func Test_Deduplicate(t *testing.T) {
	// string
	ss := []string{"AA", "aa", "Aa", "ba", "ca", "aa"}
	t.Log(Deduplicate(ss).([]string))
	// [AA aa Aa ba ca]
	// int
	is := []int{12, 32, 12, 32, 121, 342, 23}
	t.Log(Deduplicate(is).([]int))
	// [12 32 121 342 23]
}
