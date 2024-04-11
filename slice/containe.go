package slice

import "strings"

// 查找sarr中，包含有fuzzys数据的元素
// 模糊搜索
func GetContainsFuzzy(sarr, fuzzys []string) []string {
	res := make([]string, 0)
	for _, e := range sarr {
		for _, f := range fuzzys {
			if !strings.Contains(e, f) {
				break
			}
		}
		res = append(res, e)
	}

	return res
}
