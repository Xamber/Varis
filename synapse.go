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

func createSynapse(in Neuron, out Neuron) {
	syn := &synapse{
		weight: rand.Float64(),
		in:     make(chan float64),
		out:    make(chan float64),
	}

	in.addOutputSynapse(syn)
	out.addInputSynapse(syn)

	go func() {
		for {
			syn.cache = <-syn.in
			outputValue := syn.cache * syn.weight
			syn.out <- outputValue
		}
	}()
}
