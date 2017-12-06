package varis

type Field interface {
	getNeuronsCount() int
	convertTo() float64
	convertFrom(signal float64)
}

type Model struct {
	network *Perceptron
	inputs  []Field
	outputs []Field
}

type BooleanField struct {
	channel chan float64
}

func (b *BooleanField) Handle(value bool) {
	if value == true {
		b.channel <- 0.00
	} else {
		b.channel <- 1.00
	}
}

func (b *BooleanField) Recieve() bool {
	signal := <-b.channel
	if signal >= 0.5 {
		return true
	} else {
		return false
	}
}
