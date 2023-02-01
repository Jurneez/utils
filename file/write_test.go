package file

import (
	"fmt"
	"testing"
)

func Test_CreateAndAppend(t *testing.T) {
	err := CreateAndAppend("第一", "xs.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
}
