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
	for l := len(network.Layers) - 1; l > 0; l-- {
		nextLayerDelta := 0.00
		for i, n := range network.Layers[l].getNeurons() {
			var neuronDelta float64
			if l == len(network.Layers)-1 {
				neuronDelta = (expected[i] - results[i]) * n.deactivation()
			} else {
				neuronDelta = layerDelta * n.deactivation()
			}
			neuronDelta *= float64(speed)
			nextLayerDelta += neuronDelta
			n.changeWeight(neuronDelta)
		}
		layerDelta = nextLayerDelta
	}

}
