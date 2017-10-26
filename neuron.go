package varis

// Neuroner interface with base biology neuron format.
type Neuroner interface {
	getConnection() *connection
	getUUID() string
	deactivation() float64

	changeWeight(delta float64)
	live()
}

// Standart implimentation of Neuroner.
type neuron struct {
	conn         connection
	uuid         string
	bias         float64
	cache        float64
	callbackFunc func(value float64)
}

func (n *neuron) getConnection() *connection {
	return &n.conn
}

func (n *neuron) getUUID() string {
	return n.uuid
}

func (n *neuron) deactivation() float64 {
	return DEACTIVATION(n.cache)
}

func (n *neuron) changeWeight(neuronDelta float64) {
	n.bias += neuronDelta
	n.getConnection().changeWeight(neuronDelta)
}

func (n *neuron) live() {
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
