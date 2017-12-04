package varis

// Dataset - simple structure for store input and expected values.
type Dataset []struct {
	Input    []float64
	Expected []float64
}

// BackPropagation train NN by input dataset for 'times' times.
func BackPropagation(network *Network, dataset Dataset, times int) {
	var neuronDelta float64

	for times > 0 {
		for _, f := range dataset {
			results := network.Calculate(f.Input)
			layerDelta := 0.0
			for l := len(network.Layers) - 1; l > 0; l-- {
				nextLayerDelta := 0.00
				for i, n := range network.Layers[l] {
					if l == len(network.Layers)-1 {
						neuronDelta = (f.Expected[i] - results[i]) * n.deactivation()
					} else {
						neuronDelta = layerDelta * n.deactivation()
					}
					neuronDelta *= float64(1)
					nextLayerDelta += neuronDelta
					n.changeWeight(neuronDelta)
				}
				layerDelta = nextLayerDelta
			}
		}
		times--
	}
}
