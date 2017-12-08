package varis

type Neuron interface {
	live()

	getConnection() *connection
	getWeight() float64
	getCache() float64

	changeWeight(neuronDelta float64)
}

// CoreNeuron - entity with float64 weight (it is bias) and cache.
// Activation result store in cache for training.
type CoreNeuron struct {
	conn   connection
	weight float64
	cache  float64
}

// changeWeight - change weight of CoreNeuron and change weight for all related synapses.
func (n *CoreNeuron) changeWeight(neuronDelta float64) {
	n.weight += neuronDelta
	n.conn.changeWeight(neuronDelta)
}

// getWeight - get weight from CoreNeuron
func (n *CoreNeuron) getWeight() float64 {
	return n.weight
}

// getWeight - get weight from CoreNeuron
func (n *CoreNeuron) getConnection() *connection {
	return &n.conn
}

// getCache - get cache from CoreNeuron
func (n *CoreNeuron) getCache() float64 {
	return n.cache
}

type inputNeuron struct {
	CoreNeuron
	connectTo chan float64
}

func INeuron(weight float64, connectTo chan float64) Neuron {
	return &inputNeuron{
		CoreNeuron: CoreNeuron{weight: weight},
		connectTo:  connectTo,
	}
}

func (neuron *inputNeuron) live() {
	for {
		neuron.conn.broadcastSignals(<-neuron.connectTo)
	}
}

type hiddenNeuron struct {
	CoreNeuron
}

func HNeuron(weight float64) Neuron {
	return &hiddenNeuron{
		CoreNeuron: CoreNeuron{weight: weight},
	}
}

func (neuron *hiddenNeuron) live() {
	for {
		vector := neuron.conn.collectSignals()
		neuron.cache = vector.sum() + neuron.weight
		neuron.conn.broadcastSignals(ACTIVATION(neuron.cache))
	}
}

type outputNeuron struct {
	CoreNeuron
	connectTo chan float64
}

func ONeuron(weight float64, connectTo chan float64) Neuron {
	return &outputNeuron{
		CoreNeuron: CoreNeuron{weight: weight},
		connectTo:  connectTo,
	}
}

func (neuron *outputNeuron) live() {
	for {
		vector := neuron.conn.collectSignals()
		neuron.cache = vector.sum() + neuron.weight
		neuron.connectTo <- ACTIVATION(neuron.cache)
	}
}
