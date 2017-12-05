package varis

type synapse struct {
	weight    float64
	in        chan float64
	out       chan float64
	cache     float64
	inNeuron  *Neuron
	outNeuron *Neuron
}

func (syn *synapse) live() {
	for {
		syn.cache = <-syn.in
		outputValue := syn.cache * syn.weight
		syn.out <- outputValue
	}
}

func ConnectNeurons(in *Neuron, out *Neuron, weight float64) {
	syn := &synapse{
		weight:    weight,
		in:        make(chan float64),
		out:       make(chan float64),
		inNeuron:  in,
		outNeuron: out,
	}

	in.conn.addOutputSynapse(syn)
	out.conn.addInputSynapse(syn)

	go syn.live()
}
