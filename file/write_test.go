package file

import (
	"fmt"
	"testing"
)

func Test_CreateAndAppend(t *testing.T) {
	err := CreateAndAppend("第一", "xs.txt", false)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Test_CreateAndAppend_NewLine(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := CreateAndAppend("第一", "xs.txt", true)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
