package varis

type Neuroner interface {
	getConnection() *connection
	getCache() float64

	changeWeight(delta float64)
	live()
}

type Neuron struct {
	conn  connection
	bias  float64
	cache float64
	aFunc func(value float64) float64
	cFunc func(value float64)
}

func (n *Neuron) getConnection() *connection {
	return &n.conn
}

func (n *Neuron) getCache() float64 {
	return n.cache
}

func (n *Neuron) changeWeight(neuronDelta float64) {
	n.bias += neuronDelta
	n.getConnection().changeWeight(neuronDelta)
}

func (n *Neuron) live() {

	if n.aFunc == nil || n.cFunc == nil {
		return
	}

	var signals []float64
	for {
		signals = n.conn.collectSignals()
		n.cache = sum(signals) + n.bias
		output := n.aFunc(n.cache)
		n.cFunc(output)
	}
}
