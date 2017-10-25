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

func CreateLayer() *layer {
	return &layer{}
}

func ConnectLayers(now Layerer, next Layerer) {
	for i := range now.getNeurons() {
		for o := range next.getNeurons() {
			createSynapse(now.getNeuronByIndex(i), next.getNeuronByIndex(o))
		}
	}
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
