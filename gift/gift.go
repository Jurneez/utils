package gift

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type Gifts struct {
	data []*Gift
}

func NewGifts() *Gifts {
	return &Gifts{
		data: make([]*Gift, 0),
	}
}

type Gift struct {
	name   string
	weight int64
	WeightRange
}

type WeightRange struct {
	min int64
	max int64
}

func (n *Gifts) AddGift(name string, weight int64) {
	if n.data == nil {
		n.data = make([]*Gift, 0)
	}

	n.data = append(n.data, &Gift{
		name:   name,
		weight: weight,
	})
}

func (n *Gifts) AssignWeight() {
	sumWeight := int64(0)
	for _, e := range n.data {
		e.min = sumWeight
		sumWeight += e.weight
		e.max = sumWeight
	}
}

func (n *Gifts) RandGet() *Gift {
	sumWeight := int64(0)
	for _, e := range n.data {
		sumWeight += e.weight
	}
	nx, _ := rand.Int(rand.Reader, big.NewInt(int64(sumWeight)))
	xs := nx.Int64()
	for _, e := range n.data {
		if xs >= e.min && xs < e.max {
			return e
		}
	}
	return nil
}

func (n *Gifts) Range() {
	for _, e := range n.data {
		fmt.Printf("name: %s,weight:%d,min:%d,max: %d \n", e.name, e.weight, e.min, e.max)
	}
}
