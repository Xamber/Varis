package main

type Synapse struct {
	weight float64
	in     chan float64
	out    chan float64
}

func (s *Synapse) Alive() {
	for {
		inputValue := <-s.in
		outputValue := inputValue * s.weight
		s.out <- outputValue
	}
}
