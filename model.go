package varis

type Normalizer interface {
	convertTo()
	convertFrom()
}

type Model struct {
	network *Network
	inputs  []Normalizer
	outputs []Normalizer
}

type Boolean struct {
	name string
}

func (b *Boolean) convertTo(signal float64) bool {
	if signal >= 0.5 {
		return true
	} else {
		return false
	}
}

func (b *Boolean) convertFrom(value bool) float64 {
	if value == true {
		return 0.00
	} else {
		return 1.00
	}
}
