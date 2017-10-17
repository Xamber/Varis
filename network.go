package main

type Network struct {
	layers []Layerer
}

func CreateNetwork(layers ...int) Network {

	var network Network

	for index, neurons := range layers {
		var layer Layer

		for i := 0; i < neurons; i++ {
			var neuron Neuron

			switch index {
			case 0:
				neuron = CreateInputNeuron()
			case len(layers) - 1:
				neuron = CreateOutputNeuron()
			default:
				neuron = CreateHiddenNeuron()
			}
			layer.neurons = append(layer.neurons, Neuron(neuron))
		}
		network.layers = append(network.layers, Layerer(&layer))
	}

	for l := 0; l < len(network.layers)-1; l++ {
		now := network.layers[l]
		next := network.layers[l+1]
		ConnectLayers(now, next)
	}

	network.RunAllNeuron()

	return network
}

func (n *Network) GetInputLayer() Layerer {
	return n.layers[0]
}

func (n *Network) GetOutputLayer() Layerer {
	return n.layers[len(n.layers)-1]
}

func (n *Network) RunAllNeuron() {
	for _, l := range n.layers {
		for _, neuron := range l.GetNeurons() {
			go neuron.Alive()
		}
	}
}

func (n *Network) Calculate(input ...float64) []float64 {

	if len(input) != n.GetInputLayer().GetCountOfNeurons() {
		panic("Check count of input value")
	}

	for i, n := range n.GetInputLayer().GetNeurons() {
		n.Handle(input[i])
	}

	output := make([]float64, n.GetOutputLayer().GetCountOfNeurons())

	for i := range output {
		outputNeuron := n.GetOutputLayer().GetNeuronByIndex(i)
		redirectable := outputNeuron.(Redirectable)
		output[i] = <-redirectable.GetOutput()
	}

	return output
}

func (n *Network) ShowStatistic() {
	for _, layer := range n.layers {
		layer.PrintInfo()
	}
}
