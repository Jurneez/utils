package reflect_struct

import (
	"testing"
	"tool/goutils/common"
)

func Test_simpleTraverse(t *testing.T) {
	simpleTraverse(common.P1)
	simpleTraverse(common.S1)
}

// func Test_traverse(t *testing.T) {
// 	traverse(common.P1)
// 	traverse(common.S1)
// }
