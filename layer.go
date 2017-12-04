package varis

// Type layer collect Neuroner
type Layer []*Neuron

func (l *Layer) AddNeuron(neuron *Neuron) {
	*l = append(*l, neuron)
}
