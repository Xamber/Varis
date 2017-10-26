package varis

import (
	"math/rand"
)

type synapse struct {
	weight float64
	uuid   string
	in     chan float64
	out    chan float64
	cache  float64
}

func (syn *synapse) live() {
	for {
		syn.cache = <-syn.in
		outputValue := syn.cache * syn.weight
		syn.out <- outputValue
	}
}

func createSynapse(in Neuroner, out Neuroner) {
	syn := &synapse{
		weight: rand.Float64(),
		uuid:   generate_uuid(),
		in:     make(chan float64),
		out:    make(chan float64),
	}

	in.getConnection().addOutputSynapse(syn)
	out.getConnection().addInputSynapse(syn)

	go syn.live()
}
