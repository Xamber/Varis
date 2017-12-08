package varis

import "sync"

// ConnectNeurons connect two neurons.
// It creates synapse and add connection to input and output CoreNeuron.
// Have weight.
func ConnectNeurons(in Neuron, out Neuron, weight float64) {
	syn := &synapse{
		weight:    weight,
		in:        make(chan float64),
		out:       make(chan float64),
		inNeuron:  in,
		outNeuron: out,
	}

	in.getConnection().addOutputSynapse(syn)
	out.getConnection().addInputSynapse(syn)

	go syn.live()
}

type synapse struct {
	weight    float64
	in        chan float64
	out       chan float64
	cache     float64
	inNeuron  Neuron
	outNeuron Neuron
}

// ConnectNeurons live is function for goroutine.
func (syn *synapse) live() {
	for {
		syn.cache = <-syn.in
		outputValue := syn.cache * syn.weight
		syn.out <- outputValue
	}
}

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

func (c *connection) broadcastSignals(value float64) {
	for _, o := range c.outSynapses {
		o.in <- value
	}
}

func (c *connection) changeWeight(delta float64) {
	for _, s := range c.inSynapses {
		s.weight += s.cache * delta
	}
}
