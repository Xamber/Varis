package varis
//
//import "fmt"
//
//type Field interface {
//	getNeuronsCount() int
//	convertTo() float64
//	convertFrom(signal float64)
//}
//
//type Model struct {
//	network *Network
//	inputs  []Field
//	outputs []Field
//}
//
//func (m *Model) Run() {
//	normalizedInput := []float64{}
//	for _, v := range m.inputs {
//		normalizedInput = append(normalizedInput, v.convertTo())
//	}
//
//	out := m.network.Calculate(normalizedInput...)
//
//	for i := range m.outputs {
//		m.outputs[i].convertFrom(out[i])
//	}
//
//	fmt.Println(out)
//}
//
//type BooleanField struct {}
//
//func (b *BooleanField) getNeuronsCount() int {
//	return 1
//}
//
//func (b *BooleanField) convertFrom(signal float64) {
//	if signal >= 0.5 {
//		*b = true
//	} else {
//		*b = false
//	}
//}
//
//func (b *BooleanField) convertTo() float64 {
//	if *b == true {
//		return 0.00
//	} else {
//		return 1.00
//	}
//}