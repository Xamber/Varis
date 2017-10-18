package main

type Frame struct {
	inputs   []float64
	expected []float64
}

type Dataset []Frame

type Trainer struct {
	network   *Network
	trainFunc func(*Network, []float64, []float64, int)
}

func (t *Trainer) setTrainFunc(newTrainFunction func(*Network, []float64, []float64, int)) {
	t.trainFunc = newTrainFunction
}

func (t *Trainer) TrainByDataset(dataset Dataset, times int) {
	for times > 0 {
		for _, f := range dataset {
			t.trainFunc(t.network, f.inputs, f.expected, 1)
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
