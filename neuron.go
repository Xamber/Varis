package main

type BaseNeuron struct {
	inSynapses  []*Synapse
	outSynapses []*Synapse
}

type Neuroner interface {
	AddOutputSynapse(syn *Synapse)
	AddInputSynapse(syn *Synapse)
}

type Neuron struct {
	bias   float64
	output chan float64
	BaseNeuron
}

func (n *BaseNeuron) AddOutputSynapse(syn *Synapse) {
	n.outSynapses = append(n.outSynapses, syn)
}

func (n *BaseNeuron) AddInputSynapse(syn *Synapse) {
	n.inSynapses = append(n.inSynapses, syn)
}

func (n *Neuron) Handle(value float64) {
	n.Broadcast(value)
}

func (n *Neuron) Broadcast(value float64) {
	for o := range n.outSynapses {
		n.outSynapses[o].in <- value
	}
}

func (n *Neuron) CollectSignals() []float64 {

	inputSignals := make([]float64, len(n.inSynapses))

	for i := range inputSignals {
		inputSignals[i] = <-n.inSynapses[i].out
	}

	return inputSignals
}

func (n *Neuron) Alive() {

	for {

		if len(n.inSynapses) == 0 {
			break
		}

		inputSignals := n.CollectSignals()
		value := sum(inputSignals) + n.bias
		outputSignal := activation_sigmoid(value)

		if len(n.outSynapses) == 0 {
			n.output <- outputSignal
			//fmt.Println(outputSignal)
		} else {
			n.Broadcast(outputSignal)
		}
	}

}
