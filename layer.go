package varis

// Type layer collect Neuroner
type Layer []Neuroner

func (l *Layer) AddNeuron(neuron Neuroner) {
	*l = append(*l, neuron)
}
