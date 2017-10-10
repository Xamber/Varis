package main

type Network struct {
	inputLayer Layer
	hiddenLayers []Layer
	outputLayer Layer
}

func CreateNetwork(layers ...int) Network {

	var network Network

	for index, neurons := range layers {

		switch index {
		case 0:
			network.inputLayer = CreateInputLayer(neurons)
		case layers[len(layers)-1]:
			network.outputLayer = CreateOutputLayer(neurons)
		default:
			network.hiddenLayers = append(network.hiddenLayers, CreateHiddenLayer(neurons))
		}
	}

	//for l := 0; l < len(network.layers)-1; l++ {
	//
	//	now := *network.layers[l].GetNeurons()
	//	next := *network.layers[l+1].GetNeurons()
	//
	//	for i := range now {
	//		for o := range next {
	//			CreateSynapse(now[i], next[o])
	//		}
	//	}
	//}
	//
	//for l := 0; l < len(network.layers); l++ {
	//	for n := 0; n < len(network.layers[l].neurons); n++ {
	//		go network.layers[l].neurons[n].Alive()
	//	}
	//}

	return network
}

func (n Network) ShowStatistic() {

	n.inputLayer.PrintInfo()

	for _, layer := range n.hiddenLayers {
		layer.PrintInfo()
	}

	n.outputLayer.PrintInfo()
}
