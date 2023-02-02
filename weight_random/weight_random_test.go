package weight_random

import (
	"fmt"
	"testing"
)

func Test_gist(t *testing.T) {
	nn := NewWeightRandrom()
	nn.AddWeightConfig("80", 80)
	nn.AddWeightConfig("3", 3)
	nn.AddWeightConfig("10", 10)
	nn.AddWeightConfig("7", 7)
	nn.AssignWeight()
	nn.Range()

	xsm := make(map[string]int)
	for i := 0; i < 100000; i++ {
		tmp := nn.RandGet()
		_, ok := xsm[tmp.name]
		if !ok {
			xsm[tmp.name] = 0
		}
		xsm[tmp.name]++
	}

	fmt.Printf("%+v \n", xsm)
}
