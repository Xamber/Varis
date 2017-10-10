package main

import (
	"math/rand"
)

func CreateNetwork(layers ...int) Network {

	var network Network

	for _, neurons := range layers {

		layer := Layer{}
		for i := 0; i < neurons; i++ {
			layer.neurons = append(layer.neurons, Neuroner(CreateNeuron()))
		}

		network.layers = append(network.layers, layer)
	}

	for l := 0; l < len(network.layers)-1; l++ {

		now := *network.layers[l].GetNeurons()
		next := *network.layers[l+1].GetNeurons()

		for i := range now {
			for o := range next {
				CreateSynapse(now[i], next[o])
			}
		}
	}

	for l := 0; l < len(network.layers); l++ {
		for n := 0; n < len(network.layers[l].neurons); n++ {
			go network.layers[l].neurons[n].Alive()
		}
	}

	return network

}

func CreateNeuron() *Neuron {
	neuron := Neuron{bias: rand.Float64(), output:make(chan float64)}
	return &neuron
}

func CreateSynapse(in Neuroner, out Neuroner) {

	syn := Synapse{
		weight: rand.Float64(),
		in:     make(chan float64),
		out:    make(chan float64),
	}

	in.AddOutputSynapse(&syn)
	out.AddInputSynapse(&syn)

	go syn.Alive()
}
