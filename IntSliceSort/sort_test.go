package intslicesort

import "testing"

func Test_InsertSort(t *testing.T) {
	arr := []int{9, 3, 1, 4, 2, 7, 8, 6, 5}
	InsertSort(arr)
	t.Log(arr)
}

//	func Test_RadixSort(t *testing.T) {
//		arr := []int64{123, 234, 564, 765, 872, 324, 654, 874, 432}
//		t.Log(RadixSort(arr))
//	}
//
//	func Test_CountSort(t *testing.T) {
//		arr := []int64{9, 4, 2, 1, 3, 2, 3, 2, 7}
//		t.Log(CountSort(arr))
//	}
// func Test_HeapSort(t *testing.T) {
// 	// arr := []int64{6, 3, 4, 8, 7, 2, 5, 10, 9, 7}
// 	// HeapSort(arr)

// 	t.Log(int(math.Pow(2, float64(1-1))))
// }

// func Test_BucketSort(t *testing.T) {
// 	arr := []int{-7, 51, 3, 121, -3, 32, 21, 43, 4, 25, 56, 77, 16, 22, 87, 56, -10, 68, 99, 70}
// 	t.Log(BucketSort(arr))
// }
