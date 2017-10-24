package varis

type Neuron interface {
	getConnection() *connection
	getCache() float64

	train(delta float64)
	live()
}

type baseNeuron struct {
	conn connection
	bias  float64
	cache float64
}

func (n *baseNeuron) getConnection() *connection {
	return &n.conn
}

func (n *baseNeuron) getCache() float64 {
	return n.cache
}

func (n *baseNeuron) train(neuronDelta float64) {
	n.bias += neuronDelta
	n.getConnection().changeWeight(neuronDelta)
}

type inputNeuron struct {
	baseNeuron
}

func (n *inputNeuron) live() {}

type hiddenNeuron struct {
	baseNeuron
}

func (n *hiddenNeuron) live() {
	var signals []float64
	for {
		signals = n.getConnection().collectSignals()
		n.cache = sum(signals) + n.bias
		output := ACTIVATION_FUNCTION(n.cache)
		n.getConnection().broadcastSignals(output)
	}
}

type outputNeuron struct {
	baseNeuron
	output chan float64
}

func (n *outputNeuron) live() {
	var signals []float64
	for {
		signals = n.getConnection().collectSignals()
		n.cache = sum(signals) + n.bias
		output := ACTIVATION_FUNCTION(n.cache)
		n.output <- output
	}
}
