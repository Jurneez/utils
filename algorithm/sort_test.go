package algorithm

import (
	"testing"
)

//	func Test_RadixSort(t *testing.T) {
//		arr := []int64{123, 234, 564, 765, 872, 324, 654, 874, 432}
//		t.Log(RadixSort(arr))
//	}
func Test_CountSort(t *testing.T) {
	arr := []int64{9, 4, 2, 1, 3, 2, 3, 2, 7}
	t.Log(CountSort(arr))
}
