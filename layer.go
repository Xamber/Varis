package varis

type Layerer interface {
	getNeurons() []Neuron
	getCountOfNeurons() int
	getNeuronByIndex(index int) Neuron
}

type layer struct {
	neurons []Neuron
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

func (l *layer) addNeuron(neuron Neuron) {
	l.neurons = append(l.neurons, neuron)
}

func (l *layer) getNeurons() []Neuron {
	return l.neurons
}

func (l *layer) getCountOfNeurons() int {
	return len(l.neurons)
}

func (l *layer) getNeuronByIndex(index int) Neuron {
	return l.neurons[index]
}
