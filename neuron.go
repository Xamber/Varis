package varis

// Standart implimentation of Neuron.
type Neuron struct {
	conn         connection
	weight       float64
	cache        float64
	collectFunc  func() []float64
	callbackFunc func(value float64)
}

func (n *Neuron) deactivation() float64 {
	return DEACTIVATION(n.cache)
}

func (n *Neuron) changeWeight(neuronDelta float64) {
	n.weight += neuronDelta
	n.conn.changeWeight(neuronDelta)
}

func (n *Neuron) live() {
	for {
		signals := n.collectFunc()

		if n.callbackFunc == nil {
			continue
		}

		n.cache = sum(signals) + n.weight
		output := ACTIVATION(n.cache)
		n.callbackFunc(output)
	}
}
