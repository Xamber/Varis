package main

type Frame struct {
	inputs   []float64
	expected []float64
}

type Dataset []Frame

type TrainFunction func(network *Network, inputs []float64, expected []float64, speed int)

func TrainByDataset(network *Network, trainFunc TrainFunction, dataset Dataset, times int) {
	repeat(func() {
		for _, f := range dataset {
			trainFunc(network, f.inputs, f.expected, 1)
		}
	}, times)
}

func BackPropagation(network *Network, inputs []float64, expected []float64, speed int) {

	results := network.Calculate(inputs...)

	layerDelta := 0.0

	for neuronIndex, n := range network.GetOutputLayer().GetNeurons() {
		delta := expected[neuronIndex] - results[neuronIndex]

		neuronDelta := delta * n.Deactivation()
		layerDelta += neuronDelta

		n.Train(neuronDelta)
	}

	for layerIndex := len(network.layers) - 2; layerIndex > 0; layerIndex-- {
		nextDelta := 0.00
		for _, n := range network.layers[layerIndex].GetNeurons() {
			neuronDelta := layerDelta * n.Deactivation()
			nextDelta += neuronDelta
			n.Train(neuronDelta)
		}
		layerDelta = nextDelta
	}

}
