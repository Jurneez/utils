package gift

import (
	"fmt"
	"testing"
)

func Test_gist(t *testing.T) {
	nn := NewGifts()
	nn.AddGift("80", 80)
	nn.AddGift("3", 3)
	nn.AddGift("10", 10)
	nn.AddGift("7", 7)
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
