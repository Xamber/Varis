package varis

// Standart implimentation of Neuroner.
type Neuron struct {
	conn         connection
	uuid         string
	bias         float64
	cache        float64
	callbackFunc func(value float64)
}

func (n *Neuron) getConnection() *connection {
	return &n.conn
}

func (n *Neuron) getUUID() string {
	return n.uuid
}

func (n *Neuron) getWeight() float64 {
	return n.bias
}

func (n *Neuron) deactivation() float64 {
	return DEACTIVATION(n.cache)
}

func (n *Neuron) changeWeight(neuronDelta float64) {
	n.bias += neuronDelta
	n.getConnection().changeWeight(neuronDelta)
}

func (n *Neuron) live() {
	if n.callbackFunc == nil {
		return
	}

	for {
		signals := n.conn.collectSignals()
		n.cache = sum(signals) + n.bias
		output := ACTIVATION(n.cache)
		n.callbackFunc(output)
	}
}
