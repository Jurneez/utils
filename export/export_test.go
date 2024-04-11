package export

import (
	"fmt"
	"testing"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Test_Excel(t *testing.T) {
	u1 := &UserInfo{Name: "q1mi", Age: 18}
	u2 := &UserInfo{Name: "q1m12i", Age: 143}
	us := make([]*UserInfo, 0)
	us = append(us, u1)
	us = append(us, u2)

	err := ExportExcel([]string{"name", "age"}, us, "")
	if err != nil {
		fmt.Printf("%+v \n", err.Error())
	}
}
