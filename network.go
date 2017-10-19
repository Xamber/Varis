package main

type Network struct {
	layers []Layerer
	output []chan float64
}

func CreateNetwork(layers ...int) Network {

	network := Network{output: make([]chan float64, 0)}

	inputLayerIndex := 0
	outputLayerIndex := len(layers) - 1

	for index, neurons := range layers {

		layer := CreateLayer()

		for i := 0; i < neurons; i++ {
			var neuron Neuron

			switch index {
			case inputLayerIndex:
				neuron = CreateInputNeuron()
			case outputLayerIndex:
				outputChan := make(chan float64)
				neuron = CreateOutputNeuron(outputChan)
				network.output = append(network.output, outputChan)
			default:
				neuron = CreateHiddenNeuron()
			}

			layer.AddNeuron(neuron)

		}
		network.AddLayer(layer)
	}

	for l := 0; l < len(network.layers)-1; l++ {
		now := network.layers[l]
		next := network.layers[l+1]
		ConnectLayers(now, next)
	}

	network.RunAllNeuron()

	return network
}

func (n *Network) AddLayer(layer Layerer) {
	n.layers = append(n.layers, Layerer(layer))
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

	output := make([]float64, len(n.output))

	for i := range output {
		output[i] = <-n.output[i]
	}

	return output
}

func (n *Network) ShowStatistic() {
	for _, layer := range n.layers {
		layer.PrintInfo()
	}
}
