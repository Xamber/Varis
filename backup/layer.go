package main

type Layer struct {
	neurons []Neuroner
}

func (l *Layer) GetNeurons() *[]Neuroner {
	return &l.neurons
}
