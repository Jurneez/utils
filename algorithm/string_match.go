package algorithm

import "fmt"

func StringMactchBruteForce(str, subStr string) int {
	if len(str) < len(subStr) { // 匹配串长度一定要小于等于目标串，否则一定匹配不上，直接 return -1
		return -1
	}

	for i := 0; i < len(str); i++ {
		j := i
		for k := 0; k < len(subStr); {
			if str[j] == subStr[k] {
				j++
				k++
			} else {
				break // 匹配失败，直接下一轮
			}
			if k == len(subStr) { // 表示匹配成功
				return j - k
			}
			if j >= len(str) { // 走到这里，表示目标串长度走到底了也没有匹配成功，且没必要break
				return -1
			}

		}
	}

	return -1
}

func StringMactchKMP(str, subStr string) int {
	nexts := GetNextArr(subStr)
	fmt.Println(nexts)
	i, j := 0, 0
	for i < len(str) {
		if str[i] == subStr[j] {
			i++
			j++
		} else {
			if j > 0 {
				j = nexts[j-1]
			} else {
				i++
			}
		}
		if j == len(subStr) {
			break
		}
	}
	if j != len(subStr) {
		return -1
	} else {
		return i - j
	}
}

func GetNextArr(subStr string) []int {
	len := len(subStr)       // 字符串总长度
	next_index := 0          // 下一次要跳到的位置的下标
	cur_index := 1           // 当前所在位置的下标
	next := make([]int, len) // 存储 下一次要跳转位置 的数组

	for cur_index < len {
		if subStr[cur_index] == subStr[next_index] {
			next_index++
			next[cur_index] = next_index
			cur_index++
		} else {
			if next_index == 0 {
				next[cur_index] = 0
				cur_index++
			} else {
				next_index = next[next_index-1]
			}
		}
	}

	return next
}
