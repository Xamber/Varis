package varis

type Field interface {
	convertTo()
	convertFrom()
}

type Model struct {
	network *Network
	inputs  []Field
	outputs []Field
}

type BooleanField struct {
	name string
}

func (b *BooleanField) convertTo(signal float64) bool {
	if signal >= 0.5 {
		return true
	} else {
		return false
	}
}

func (b *BooleanField) convertFrom(value bool) float64 {
	if value == true {
		return 0.00
	} else {
		return 1.00
	}
}
