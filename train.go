package varis

// Dataset - simple type for store input and expected Vectors.
type Dataset [][2]Vector

// PerceptronTrainer is a trainer for Perceptron networks
type PerceptronTrainer struct {
	Network *Perceptron
	Dataset Dataset
}

// BackPropagation train Network input Dataset for 'times' times.
func (t *PerceptronTrainer) BackPropagation(times int) {
	var neuronDelta float64

	for iteration := 0; iteration < times; iteration++ {
		for _, frame := range t.Dataset {
			expected := frame[1]
			results := t.Network.Calculate(frame[0])

			t.Network.mux.Lock()

			layerDelta := 0.0
			for l := len(t.Network.layers) - 1; l > 0; l-- {
				nextLayerDelta := 0.00
				for i, n := range t.Network.layers[l] {
					if l == len(t.Network.layers)-1 {
						neuronDelta = (expected[i] - results[i]) * DEACTIVATION(n.getCore().cache)
					} else {
						neuronDelta = layerDelta * DEACTIVATION(n.getCore().cache)
					}
					neuronDelta *= float64(1)
					nextLayerDelta += neuronDelta
					n.changeWeight(neuronDelta)
				}
				layerDelta = nextLayerDelta
			}

			t.Network.mux.Unlock()
		}
	}
}
