package varis

// Type layer collect Neuron
type Layer []*Neuron

func (l *Layer) AddNeuron(neuron *Neuron) {
	*l = append(*l, neuron)
}
