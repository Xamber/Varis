package main

type Network struct {
	inputLayer   Layerer
	hiddenLayers []Layerer
	outputLayer  Layerer
}

func CreateNetwork(layers ...int) Network {

	var network Network

	for index, neurons := range layers {
		switch index {
		case 0:
			network.inputLayer = CreateInputLayer(neurons)
		case len(layers)-1:
			network.outputLayer = CreateOutputLayer(neurons)
		default:
			network.hiddenLayers = append(network.hiddenLayers, CreateHiddenLayer(neurons))
		}
	}

	for l := 0; l < len(network.hiddenLayers)-1; l++ {
		now := network.hiddenLayers[l]
		next := network.hiddenLayers[l+1]
		ConnectLayers(now, next)
	}

	ConnectLayers(network.inputLayer, network.hiddenLayers[0])
	ConnectLayers(network.hiddenLayers[len(network.hiddenLayers)-1], network.outputLayer)

	for l := 0; l < len(network.hiddenLayers); l++ {
		network.hiddenLayers[l].RunAllNeurons()
	}

	network.outputLayer.RunAllNeurons()

	return network
}

func (n *Network) Calculate(input []float64) []float64 {

	if len(input) != n.inputLayer.GetCountOfNeurons() {
		panic("Check count of input value")
	}

	for i, n := range n.inputLayer.GetNeurons() {
		n.Handle(input[i])
	}

	output := make([]float64, n.outputLayer.GetCountOfNeurons())
	for i := range output {
		output[i] = <-n.outputLayer.GetNeuronByIndex(i).(RedirectNeuron).GetOutput()
	}

	return output
}

func (n *Network) ShowStatistic() {

	n.inputLayer.PrintInfo()

	for _, layer := range n.hiddenLayers {
		layer.PrintInfo()
	}

	n.outputLayer.PrintInfo()
}
