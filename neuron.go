package varis

// Standart implimentation of Neuron.
type Neuron struct {
	conn           connection
	weight         float64
	cache          float64
	collectFunc    func() Vector
	activationFunc func(vector Vector) float64
	callbackFunc   func(value float64)
}

func (n *Neuron) deactivation() float64 {
	return DEACTIVATION(n.cache)
}

func (n *Neuron) changeWeight(neuronDelta float64) {
	n.weight += neuronDelta
	n.conn.changeWeight(neuronDelta)
}

func (n *Neuron) live() {

	if n.collectFunc == nil && n.activationFunc == nil && n.callbackFunc == nil {
		panic("Neuron do nothing")
	}

	if n.collectFunc == nil {
		return
	}

	for {
		signals := n.collectFunc()
		output := n.activationFunc(signals)
		n.callbackFunc(output)
	}
}

func (n *Neuron) standartActivation(vector Vector) float64 {
	n.cache = vector.sum() + n.weight
	output := ACTIVATION(n.cache)
	return output
}

func (n *Neuron) SetPipeActivation() {
	n.activationFunc = func(vector Vector) float64 {
		return vector.sum()
	}
}

func (n *Neuron) SetRedirectOutput(outputChan chan float64) {
	redirect := func(value float64) {
		outputChan <- value
	}
	n.callbackFunc = redirect
}

func (n *Neuron) SetExternalInput(outputChan chan float64) {
	redirect := func(value float64) {
		outputChan <- value
	}
	n.callbackFunc = redirect
}
