package varis

type layer []Neuroner

func (l *layer) AddNeuron(neuron Neuroner) {
	*l = append(*l, neuron)
}
