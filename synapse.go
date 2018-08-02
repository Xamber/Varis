package varis

import "sync"

// ConnectNeurons connect two neurons.
// It creates synapse and add connection to input and output Neuron.
func ConnectNeurons(in Neuron, out Neuron, weight float64) {
	syn := &synapse{
		weight:    weight,
		in:        make(chan float64),
		out:       make(chan float64),
		inNeuron:  in,
		outNeuron: out,
	}

	in.getCore().conn.addOutputSynapse(syn)
	out.getCore().conn.addInputSynapse(syn)

	go syn.live()
}

// synapse is connection between two neurons.
// Store pointers to input and output neuron.
// Have two channels for data and weight.
type synapse struct {
	weight    float64
	in        chan float64
	out       chan float64
	cache     float64
	inNeuron  Neuron
	outNeuron Neuron
}

// ConnectNeurons live is function for goroutine.
// It receive data from input. Multiply input to weight and send to output.
func (syn *synapse) live() {
	for {
		syn.cache = <-syn.in
		outputValue := syn.cache * syn.weight
		syn.out <- outputValue
	}
}

// connection store input and output synapses
type connection struct {
	inSynapses  []*synapse
	outSynapses []*synapse
}

func (c *connection) addOutputSynapse(syn *synapse) {
	c.outSynapses = append(c.outSynapses, syn)
}

func (c *connection) addInputSynapse(syn *synapse) {
	c.inSynapses = append(c.inSynapses, syn)
}

// collectSignals wait signals for all of input Synapses and return Vector with values.
func (c *connection) collectSignals() Vector {
	inputSignals := make(Vector, len(c.inSynapses))

	wg := sync.WaitGroup{}
	wg.Add(len(c.inSynapses))

	for i := range inputSignals {
		go func(index int) {
			inputSignals[index] = <-c.inSynapses[index].out
			wg.Done()
		}(i)
	}

	wg.Wait()
	return inputSignals
}

// broadcastSignals send signal to all output synapses.
func (c *connection) broadcastSignals(value float64) {
	for _, o := range c.outSynapses {
		o.in <- value
	}
}

// changeWeight is function for train.
func (c *connection) changeWeight(delta float64) {
	for _, s := range c.inSynapses {
		s.weight += s.cache * delta
	}
}
