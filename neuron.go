package varis

// Neuron - interface for all Neuron.
// Each Neuron must have:
// - coreNeuron is a basic neuron for all types
// - getCore() is a the function for getting pointer to CoreNeuron
// - live() - method for running neuron's goroutine. All kind of Neurons implement functionality live
// - changeWeight is the method for training
type Neuron interface {
	live()
	getCore() *CoreNeuron
	changeWeight(neuronDelta float64)
}

// CoreNeuron - entity with float64 weight (it is bias) and connection.
// Activation result store in cache for training.
type CoreNeuron struct {
	conn   connection
	weight float64
	cache  float64
}

// changeWeight - change weight of CoreNeuron and change weight for all related synapses (with connection.changeWeight).
func (n *CoreNeuron) changeWeight(neuronDelta float64) {
	n.weight += neuronDelta
	n.conn.changeWeight(neuronDelta)
}

// getCore - return core of Neuron.
func (n *CoreNeuron) getCore() *CoreNeuron {
	return n
}

// This kind of Neuron get signal from connectTo channel and broadcast it to all output synapses without Activation.
type inputNeuron struct {
	CoreNeuron
	connectTo chan float64
}

// INeuron - creates inputNeuron.
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

// This kind of Neuron get signal from input Synapses channel, activate and broadcast it to all output synapses.
type hiddenNeuron struct {
	CoreNeuron
}

// HNeuron - creates hiddenNeuron.
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

// This kind of Neuron get signal from input Synapses channel, activate and send it to connectTo channel.
type outputNeuron struct {
	CoreNeuron
	connectTo chan float64
}

// ONeuron - creates outputNeuron.
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
