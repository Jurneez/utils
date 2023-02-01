package print_struct

import (
	"testing"
	"tool/goutils/common"
)

func Test_PrintStructString(t *testing.T) {
	p_str := PrintStructString(common.P1)
	t.Logf("PrintStructString==>[\n%s\n] \n ", string(p_str))
}

func Test_PrintStructJson(t *testing.T) {
	p_str := PrintStructJson(common.P1)
	t.Logf("PrintStructJson==>[\n%s\n] \n ", string(p_str))
}
