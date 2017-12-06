package varis

// Dataset - simple type for store input and expected values.
type Dataset [][2]Vector

// BackPropagation train NN by input dataset for 'times' times.
func BackPropagation(network *Perceptron, dataset Dataset, times int) {
	var neuronDelta float64

	var lastLayerIndex = len(network.layers) - 1

	for iteration := 0; iteration < times; iteration++ {
		for _, frame := range dataset {
			results := network.Calculate(frame[0])
			layerDelta := 0.0
			for l := lastLayerIndex; l > 0; l-- {
				nextLayerDelta := 0.00
				for i, n := range network.layers[l] {
					if l == lastLayerIndex {
						neuronDelta = (frame[1][i] - results[i]) * DEACTIVATION(n.cache)
					} else {
						neuronDelta = layerDelta * DEACTIVATION(n.cache)
					}
					neuronDelta *= float64(1)
					nextLayerDelta += neuronDelta
					n.changeWeight(neuronDelta)
				}
				layerDelta = nextLayerDelta
			}
		}
	}
}
