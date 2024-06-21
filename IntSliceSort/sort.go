package intslicesort

import (
	"fmt"
)

// 交换排序——冒泡排序
func BubbleSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ { // 每一次for循环之后，最大的数据就放在了后面
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	return numbers
}

// 交换排序——快速排序
func QuickSort(low, high int, numbers []int) {
	if low < high {
		mid := quick_sort_get_middle(low, high, numbers)
		QuickSort(0, mid-1, numbers)
		QuickSort(mid+1, high, numbers)
	}
}

func quick_sort_get_middle(low, high int, numbers []int) int {
	if low == high {
		return 0
	}
	for low < high {
		if numbers[low] > numbers[high] {
			numbers[low], numbers[high] = numbers[high], numbers[low]
			low++
		} else {
			high--
		}
	}

	return low
}

// 选择排序
func SelectSort(numbers []int) {
	if len(numbers) == 0 {
		return
	}
	for i := 0; i < len(numbers); i++ {
		min := i                                // 保存最小值的下标
		for j := i + 1; j < len(numbers); j++ { // 找到最小值所在的下标
			if numbers[j] < numbers[min] {
				min = j
			}
		}
		if numbers[min] < numbers[i] { // 进行比较、交换
			numbers[min], numbers[i] = numbers[i], numbers[min]
		}
		fmt.Println(numbers)
	}
}

// 插入排序
func InsertSort(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		fmt.Println(numbers)
		for j := i + 1; j > 0; j-- {
			if numbers[j] < numbers[j-1] {
				numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
			} else {
				break
			}
		}
	}
	fmt.Println(numbers)
}

// 希尔排序
func ShellSort(numbers []int) {
	gap := (len(numbers) + 1) / 2
	fmt.Println("原数组：", numbers)
	for gap > 0 {
		for i := 0; i < len(numbers)-gap; i++ { // 每次循环的数组数量都翻倍，但是每个数组都基本有序
			for j := i + gap; j > gap-1; j = j - gap { // 每次和间隔gap的数据进行比较
				if numbers[j] < numbers[j-gap] {
					numbers[j], numbers[j-gap] = numbers[j-gap], numbers[j]
				} else {
					break
				}
			}
		}
		fmt.Println("gap = ", gap, " 每次排序后的数组为：", numbers)
		gap = gap / 2
	}
}

// 归并排序
func MergeSort(numbers []int) []int {
	if len(numbers) <= 1 { // 特殊情况直接返回
		return numbers
	}
	mid := len(numbers) / 2
	left_arr := MergeSort(numbers[:mid])
	right_arr := MergeSort(numbers[mid:])
	fmt.Println("每次排序的数组为：", left_arr, " 和 ", right_arr)
	res := mergeArr(left_arr, right_arr)
	fmt.Println("排序结果为：", res)
	return res
}

// numbers1,numbers2两个数组均有序
func mergeArr(numbers1, numbers2 []int) []int {
	numbers1_len, numbers2_len := len(numbers1), len(numbers2)
	res := make([]int, 0, numbers1_len+numbers2_len)

	numbers1_i, numbers2_i := 0, 0
	for numbers1_i < numbers1_len && numbers2_i < numbers2_len {
		if numbers1[numbers1_i] > numbers2[numbers2_i] {
			res = append(res, numbers2[numbers2_i]) // 比较小数据 push的res中
			numbers2_i++
		} else {
			res = append(res, numbers1[numbers1_i]) // 比较小数据 push的res中
			numbers1_i++
		}
	}

	// for循环之后，可能存在还没有push到res中的数据
	if numbers1_i == 0 {
		res = append(res, numbers1[:]...)
	} else if numbers1_i < numbers1_len {
		res = append(res, numbers1[numbers1_i:]...)
	}

	if numbers2_i == 0 {
		res = append(res, numbers2[:]...)
	} else if numbers2_i < numbers2_len {
		res = append(res, numbers2[numbers2_i:]...)
	}

	return res
}

// 堆排序：数据从小到大排
// 构建大堆顶
func HeapSort(numbers []int64) {
	length := len(numbers)
	for i := 0; i < length; i++ {
		fmt.Println("待排序的初始值", numbers[:length-i])
		buildBigHeap(numbers[:length-i], 0)
		fmt.Println("堆排序结果", numbers)
		numbers[0], numbers[length-i-1] = numbers[length-i-1], numbers[0]
		fmt.Println("排序结果", numbers)
	}
}

// 单节点数据构建
/*
len = 数组长度
叶子节点： i <= len/2+1
非叶子节点：i > len/2 +1
*/
func buildBigHeap(numbers []int64, index int64) {
	if index > int64(len(numbers)-1) {
		return
	}
	temp := len(numbers)/2 + 1
	if index > int64(temp) { // 叶子节点，不做处理
		return
	} else { // 非叶子节点
		buildBigHeap(numbers, 2*index+1) // 处理左节点
		buildBigHeap(numbers, 2*index+2) // 处理右节点
		if (2*index+1) <= int64(len(numbers)-1) && numbers[index] < numbers[2*index+1] {
			numbers[index], numbers[2*index+1] = numbers[2*index+1], numbers[index]
		}
		if (2*index+2) <= int64(len(numbers)-1) && numbers[index] < numbers[2*index+2] {
			numbers[index], numbers[2*index+2] = numbers[2*index+2], numbers[index]
		}
	}
}

// 非比较排序算法——基数排序算法：适合正整数排序
/*
时间复杂度：O(nk)
*/
func RadixSort(numbers []int64) []int64 {
	if len(numbers) == 0 {
		return numbers
	}

	// 获取数组中的最大值
	max_number := GetInt64ArrayMaxValue(numbers) // 时间复杂度：O(n)
	// 获取最大值位数
	max_number_digits := CountDigits(int(max_number))

	for m := 0; m < max_number_digits; m++ { // 时间复杂度：O(m)
		// 1. 申请10个桶资源
		barrels := make([][]int64, 10)

		// 2. 将数据放入对应桶中
		digit := int64(1)
		for k := 0; k < m; k++ {
			digit = digit * 10
		}

		for i := 0; i < len(numbers); i++ { // 遍历numbers中的数据，并放入barrels中   时间复杂度：O(n)
			digit_value, _ := GetDigitValue(numbers[i], digit)              // 获取数据指定位数的数值
			barrels[digit_value] = append(barrels[digit_value], numbers[i]) // 根据数值将数据写入对应桶中
		}
		fmt.Println("每次循环后，桶中的数据：", barrels)

		// 3. 组合数组
		numbers_tmp := make([]int64, 0, len(numbers))
		for i := 0; i < 10; i++ {
			if len(barrels[i]) > 0 {
				numbers_tmp = append(numbers_tmp, barrels[i]...)
			}
		}
		numbers = numbers_tmp
		fmt.Println("每次循环后，数组中的数据：", numbers)
	}

	return numbers
}

// 非比较排序算法——计数排序：适合数据跨度比较小的数组
func CountSort(numbers []int64) []int64 {
	if len(numbers) == 0 {
		return numbers
	}

	// 1. 找出最大值和最小值
	min, max := GetInt64ArrayMaxAndMinValue(numbers) // 时间复杂度O(n)
	// 2.申请映射数组，映射数组长度为：最大值-最小值+1。适当节省空间
	mapping_arr := make([]int64, max-min+1)

	// 3. 将待排序数值所在映射数组下标的位置加一：位置=（arr[i]-最小值）
	for i := 0; i < len(numbers); i++ {
		mapping_arr[numbers[i]-min]++
	}
	fmt.Println("遍历后的映射数组为：", mapping_arr)

	// 4. 计算每个数据在数组中所占位置的跨度，为了保证后续组合结果时的稳定性
	countArr := make([]int64, len(mapping_arr))
	for i := 0; i < len(numbers); i++ {
		if i == 0 {
			countArr[i] = mapping_arr[i]
			continue
		}
		countArr[i] = countArr[i-1] + mapping_arr[i]
	}
	fmt.Println("遍历后的count数组为：", countArr)

	res := make([]int64, len(numbers))
	// 4. 遍历映射数组，得到有序数据
	for i := 0; i < len(countArr); i++ {
		res[countArr[numbers[i]-min]-1] = numbers[i]
		countArr[numbers[i]-min]--
	}
	return res
}
