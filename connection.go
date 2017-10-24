package varis

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

func (c *connection) getOutputSynapse() []*synapse {
	return c.outSynapses
}

func (c *connection) getInputSynapse() []*synapse {
	return c.outSynapses
}

func (c *connection) collectSignals() []float64 {

	inputSignals := make([]float64, len(c.inSynapses))

	for i := range inputSignals {
		inputSignals[i] = <-c.inSynapses[i].out
	}

	return inputSignals
}

func (c *connection) broadcastSignals(value float64) {
	for _, o := range c.outSynapses {
		o.in <- value
	}
}
