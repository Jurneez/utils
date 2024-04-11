package slice

import "testing"

// func Test_Deduplicate(t *testing.T) {
// 	// string
// 	ss := []string{"AA", "aa", "Aa", "ba", "ca", "aa"}
// 	t.Log(Deduplicate(ss).([]string))
// 	// [AA aa Aa ba ca]
// 	// int
// 	is := []int{12, 32, 12, 32, 121, 342, 23}
// 	t.Log(Deduplicate(is).([]int))
// 	// [12 32 121 342 23]
// }

func Test_Intersection(t *testing.T) {
	// string
	ss := []string{"AA", "aa", "Aa", "ba", "ca", "aa"}
	ss1 := []string{"AB", "aa", "Aa", "ba", "csa", "aa"}
	t.Log(Intersection(ss, ss1).([]string))
	// [12 32 121 342 23]
}
