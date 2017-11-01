package varis
//
//import (
//	"testing"
//
//	"fmt"
//)
//
//func Run(input []Field, output []Field) {
//	network := CreateNetwork(2, 3, 1)
//
//	normalizedInput := []float64{}
//	for _, v := range input {
//		normalizedInput = append(normalizedInput, v.convertTo())
//	}
//
//	out := network.Calculate(normalizedInput...)
//
//	for i := range output {
//		output[i].convertFrom(out[i])
//	}
//
//	fmt.Println(out)
//}
//
//func BenchmarkModel(b *testing.B) {
//
//	x1 := BooleanField{}
//	x2 := &BooleanField(false)
//	o := &BooleanField(false)
//
//	m := Model{
//		inputs: []Field{x1, x2},
//		outputs: []Field{o},
//		network: &CreateNetwork(2,2,1),
//	}
//
//	x1 := BooleanField(true)
//	x2 := BooleanField(false)
//
//	o := BooleanField(true)
//
//	Run([]Field{&BooleanField(true), &x2}, []Field{&o})
//}
//
////func createModel(input []Field, output []Field, hiddenLayers int) *Model {
////
////	inputNeuronsCount := 0
////	outputNeuronsCount := 0
////
////	var networkConfig []int
////
////	for _, inputFiled := range input {
////		inputNeuronsCount += inputFiled.getNeuronsCount()
////	}
////
////	for _, outputFiled := range output {
////		outputNeuronsCount += outputFiled.getNeuronsCount()
////	}
////
////	// Rewtire me !!!!
////	networkConfig = append(networkConfig, inputNeuronsCount)
////	for hiddenLayers > 0 {
////		networkConfig = append(networkConfig, inputNeuronsCount)
////		hiddenLayers--
////	}
////	networkConfig = append(networkConfig, outputNeuronsCount)
////
////	network := CreateNetwork(networkConfig...)
////
////	start := 0
////	for i, _ := range input {
////		input[i].setNeurons(network.GetInputLayer()[start:start+input[i].getNeuronsCount()])
////	}
////
////	start = 0
////	for i, _ := range output {
////		output[i].setNeurons(network.GetOutputLayer()[start:start+output[i].getNeuronsCount()])
////	}
////
////	return &Model{network: &network, inputs: input, outputs: output}
////
////}
