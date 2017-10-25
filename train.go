package varis

type Dataset []struct {
	input    []float64
	expected []float64
}

type Trainer struct {
	Network   *Network
	TrainFunc func(*Network, []float64, []float64, int)
}

func (t *Trainer) TrainByDataset(dataset Dataset, times int) {
	for times > 0 {
		for _, f := range dataset {
			t.TrainFunc(t.Network, f.input, f.expected, 1)
		}
		times--
	}
}

func BackPropagation(network *Network, inputs []float64, expected []float64, speed int) {
	results := network.Calculate(inputs...)

	layerDelta := 0.0
	for i, n := range network.GetOutputLayer().getNeurons() {
		neuronDelta := (expected[i] - results[i]) * n.deactivation()
		neuronDelta *= float64(speed)
		layerDelta += neuronDelta
		n.changeWeight(neuronDelta)
	}

	for i := len(network.Layers) - 2; i > 0; i-- {
		nextLayerDelta := 0.00
		for _, n := range network.Layers[i].getNeurons() {
			neuronDelta := layerDelta * n.deactivation()
			nextLayerDelta += neuronDelta
			n.changeWeight(neuronDelta)
		}
		layerDelta = nextLayerDelta
	}
}
