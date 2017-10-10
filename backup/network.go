package main

import (
	"fmt"
)

type Network struct {
	layers []Layer
	output []chan float64
}

func (n *Network) Calculate(input []float64) []float64 {

	n.HandleInput(input)

	outputNeurons := *n.GetOutputNeurons()

	output := make([]float64, len(outputNeurons))

	for i := range outputNeurons {
		output[i] = <-outputNeurons[i].GetOutput()
	}

	return output
}


func (n Network) HandleInput(input []float64) {

	inputNeurons := *n.GetInputNeurons()

	if len(input) != len(inputNeurons) {
		panic("Check count of input value")
	}

	for i := range inputNeurons {
		inputNeurons[i].Handle(input[i])
	}
}

func (n Network) GetCountOfLayers() int {
	return len(n.layers)
}

func (n Network) GetInputNeurons() *[]Neuroner {
	return n.layers[0].GetNeurons()
}

func (n Network) GetOutputNeurons() *[]Neuroner {
	return n.layers[n.GetCountOfLayers()-1].GetNeurons()
}

func (n Network) ShowStatistic() {
	for _, layer := range n.layers {
		fmt.Println("Layer:")

		for index, neuron := range *layer.GetNeurons() {
			fmt.Println("    Neuron (", index, "): ", &neuron)

			for _, synapse := range neuron.GetInputSynapses() {
				fmt.Println("        InSynapse: ", synapse)
			}

			for _, synapse := range neuron.GetOutputSynapses() {
				fmt.Println("        OutSynapse: ", synapse)
			}
		}
	}
}
