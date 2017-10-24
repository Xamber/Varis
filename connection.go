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
