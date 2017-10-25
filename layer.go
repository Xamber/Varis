package varis

type Layerer interface {
	AddNeuron(neuron Neuroner)
	getNeurons() []Neuroner
	getCountOfNeurons() int
	getNeuronByIndex(index int) Neuroner
}

type layer struct {
	neurons []Neuroner
}

func (l *layer) AddNeuron(neuron Neuroner) {
	l.neurons = append(l.neurons, neuron)
}

func (l *layer) getNeurons() []Neuroner {
	return l.neurons
}

func (l *layer) getCountOfNeurons() int {
	return len(l.neurons)
}

func (l *layer) getNeuronByIndex(index int) Neuroner {
	return l.neurons[index]
}
