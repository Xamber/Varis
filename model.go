package varis

type Field interface {
	getNeuronsCount() int
	convertTo() float64
	convertFrom(signal float64)
}

type Model struct {
	network *Network
	inputs  []Field
	outputs []Field
}

type Bool bool

func (b *Bool) getNeuronsCount() int {
	return 1
}

func (b *Bool) convertFrom(signal float64) {
	if signal >= 0.5 {
		*b = true
	} else {
		*b = false
	}
}

func (b *Bool) convertTo() float64 {
	if *b == true {
		return 0.00
	} else {
		return 1.00
	}
}
