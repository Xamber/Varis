package varis

type Frame struct {
	Inputs   []float64
	Expected []float64
}

type Dataset []Frame

type Trainer struct {
	Network   *Network
	TrainFunc func(*Network, []float64, []float64, int)
}

func (t *Trainer) SetTrainFunc(newTrainFunction func(*Network, []float64, []float64, int)) {
	t.TrainFunc = newTrainFunction
}

func (t *Trainer) TrainByDataset(dataset Dataset, times int) {
	for times > 0 {
		for _, f := range dataset {
			t.TrainFunc(t.Network, f.Inputs, f.Expected, 1)
		}
		times--
	}
}

func BackPropagation(network *Network, inputs []float64, expected []float64, speed int) {

	results := network.Calculate(inputs...)

	layerDelta := 0.0

	for neuronIndex, n := range network.GetOutputLayer().GetNeurons() {
		neuronDelta := (expected[neuronIndex] - results[neuronIndex]) * n.Deactivation() * float64(speed)
		layerDelta += neuronDelta
		n.Train(neuronDelta)
	}

	for layerIndex := len(network.layers) - 2; layerIndex > 0; layerIndex-- {
		nextLayerDelta := 0.00
		for _, n := range network.layers[layerIndex].GetNeurons() {
			neuronDelta := layerDelta * n.Deactivation()
			nextLayerDelta += neuronDelta
			n.Train(neuronDelta)
		}
		layerDelta = nextLayerDelta
	}

}
