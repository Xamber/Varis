package varis

import (
	"math/rand"
)

type synapse struct {
	weight float64
	in     chan float64
	out    chan float64
	cache  float64
}

func (syn *synapse) Live() {
	for {
		syn.cache = <-syn.in
		outputValue := syn.cache * syn.weight
		syn.out <- outputValue
	}
}

func createSynapse(in Neuron, out Neuron) {
	syn := &synapse{
		weight: rand.Float64(),
		in:     make(chan float64),
		out:    make(chan float64),
	}

	in.getConnection().addOutputSynapse(syn)
	out.getConnection().addInputSynapse(syn)

	go syn.Live()
}
