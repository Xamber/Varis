package varis

type Neuroner interface {
	getConnection() *connection
	deactivation() float64

	changeWeight(delta float64)
	live()
}

type Neuron struct {
	conn         connection
	bias         float64
	cache        float64
	callbackFunc func(value float64)
}

func (n *Neuron) getConnection() *connection {
	return &n.conn
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

	var signals []float64

	for {
		signals = n.conn.collectSignals()
		n.cache = sum(signals) + n.bias
		output := ACTIVATION(n.cache)
		n.callbackFunc(output)
	}
}
