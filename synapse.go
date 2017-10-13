package main

import (
	"math/rand"
)

type Synapse struct {
	weight float64
	in     chan float64
	out    chan float64
	cache  float64
}

func CreateSynapse(in Neuron, out Neuron) {
	syn := &Synapse{
		weight: rand.Float64(),
		in:     make(chan float64),
		out:    make(chan float64),
	}

	in.AddOutputSynapse(syn)
	out.AddInputSynapse(syn)

	go syn.Alive()
}

func (s *Synapse) ChangeWeight(delta float64) {
	s.weight += s.cache * delta
}

func (s *Synapse) Alive() {
	for {
		s.cache = <-s.in
		outputValue := s.cache * s.weight
		s.out <- outputValue
	}
}
