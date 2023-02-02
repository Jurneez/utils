package weight_random

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type WeightRandrom struct {
	data []*WeightConfig
}

func NewWeightRandrom() *WeightRandrom {
	return &WeightRandrom{
		data: make([]*WeightConfig, 0),
	}
}

type WeightConfig struct {
	name   string
	weight int64
	WeightRange
}

type WeightRange struct {
	min int64
	max int64
}

func (n *WeightRandrom) AddWeightConfig(name string, weight int64) {
	if n.data == nil {
		n.data = make([]*WeightConfig, 0)
	}

	n.data = append(n.data, &WeightConfig{
		name:   name,
		weight: weight,
	})
}

func (n *WeightRandrom) AssignWeight() {
	sumWeight := int64(0)
	for _, e := range n.data {
		e.min = sumWeight
		sumWeight += e.weight
		e.max = sumWeight
	}
}

func (n *WeightRandrom) RandGet() *WeightConfig {
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

func (n *WeightRandrom) Range() {
	for _, e := range n.data {
		fmt.Printf("name: %s,weight:%d,min:%d,max: %d \n", e.name, e.weight, e.min, e.max)
	}
}
