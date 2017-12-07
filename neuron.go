package varis

// Neuron type Enum
const (
	InputNeuron = iota
	HiddenNeuron
	OutputNeuron
)

// Neuron - entity with float64 weight (it is bias) and cache.
// Neuron have live method, which ran collectFunc, activationFunc and callbackFunc.
// Activation result store in cache for training.
type Neuron struct {
	conn           connection
	weight         float64
	cache          float64
	collectFunc    func() Vector
	activationFunc func(vector Vector) float64
	callbackFunc   func(value float64)
}

// CreateNeuron - create Neuron by type.
// It creates all function for live method in set its to Neuron.
// If Neuron have input or output channel - CreateNeuron return channel.
func CreateNeuron(neuronType int, weight float64) (*Neuron, chan float64) {

	var neuron = &Neuron{weight: weight}
	var channel chan float64

	neuron.callbackFunc = neuron.conn.broadcastSignals
	neuron.collectFunc = neuron.conn.collectSignals
	neuron.activationFunc = func(vector Vector) float64 {
		neuron.cache = vector.sum() + neuron.weight
		return ACTIVATION(neuron.cache)
	}

	switch neuronType {
	case InputNeuron:
		channel = make(chan float64)
		neuron.activationFunc = func(vector Vector) float64 {
			return vector.sum()
		}
		neuron.collectFunc = func() Vector {
			return Vector{<-channel}
		}
	case HiddenNeuron:
		channel = nil
	case OutputNeuron:
		channel = make(chan float64)
		neuron.callbackFunc = func(value float64) {
			channel <- value
		}
	}
	return neuron, channel
}

// changeWeight - change weight of Neuron and change weight for all related synapses.
func (n *Neuron) changeWeight(neuronDelta float64) {
	n.weight += neuronDelta
	n.conn.changeWeight(neuronDelta)
}

// live - method for goroutine.
func (n *Neuron) live() {

	if n.activationFunc == nil && n.callbackFunc == nil {
		panic("Neuron do nothing")
	}

	if n.collectFunc == nil {
		return
	}

	for {
		n.callbackFunc(n.activationFunc(n.collectFunc()))
	}
}
