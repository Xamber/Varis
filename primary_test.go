package varis

import (
	"math/rand"
	"testing"
)

var dataset = Dataset{
	{Vector{0.0, 0.0}, Vector{1.0}},
	{Vector{1.0, 0.0}, Vector{0.0}},
	{Vector{0.0, 1.0}, Vector{0.0}},
	{Vector{1.0, 1.0}, Vector{1.0}},
}

func TestNetworkCreation(t *testing.T) {
	_ = CreatePerceptron(2, 3, 1)
}

func TestNetworkCalculate(t *testing.T) {
	rand.Seed(1488)
	nn := CreatePerceptron(2, 3, 1)

	if nn.Calculate(Vector{0.0, 0.0})[0] != 0.6631222149019123 {
		t.Error("Wrong with calculation")
	}
}

func TestNetworkTrain(t *testing.T) {
	nn := CreatePerceptron(2, 3, 1)

	trainer := PerceptronTrainer{&nn, dataset}
	trainer.BackPropagation(10000)

	if nn.Calculate(Vector{0.0, 0.0})[0] < 0.800000 {
		t.Log("Bad calculation: ", Vector{0.0, 0.0}, " result: ", nn.Calculate(Vector{0.0, 0.0})[0])
	}

	if nn.Calculate(Vector{1.0, 0.0})[0] > 0.200000 {
		t.Log("Bad calculation: ", Vector{1.0, 0.0}, " result: ", nn.Calculate(Vector{1.0, 0.0})[0])
	}

	if nn.Calculate(Vector{0.0, 1.0})[0] > 0.200000 {
		t.Log("Bad calculation: ", Vector{0.0, 1.0}, " result: ", nn.Calculate(Vector{0.0, 1.0})[0])
	}

	if nn.Calculate(Vector{1.0, 1.0})[0] < 0.800000 {
		t.Log("Bad calculation: ", Vector{1.0, 1.0}, " result: ", nn.Calculate(Vector{1.0, 1.0})[0])
	}
}

func TestEncodeJson(t *testing.T) {
	n1 := CreatePerceptron(2, 3, 1)

	trainer := PerceptronTrainer{&n1, dataset}
	trainer.BackPropagation(10000)

	json := ToJSON(n1)
	n2 := FromJSON(json)

	if n1.Calculate(Vector{0.0, 0.0})[0] != n2.Calculate(Vector{0.0, 0.0})[0] {
		t.Error("Result of neural newtwork calculating not a same ")
	}

	if n1.Calculate(Vector{1.0, 0.0})[0] != n2.Calculate(Vector{1.0, 0.0})[0] {
		t.Error("Result of neural newtwork calculating not a same")
	}

	if n1.Calculate(Vector{0.0, 1.0})[0] != n2.Calculate(Vector{0.0, 1.0})[0] {
		t.Error("Result of neural newtwork calculating not a same ")
	}

	if n1.Calculate(Vector{1.0, 1.0})[0] != n2.Calculate(Vector{1.0, 1.0})[0] {
		t.Error("Result of neural newtwork calculating not a same ")
	}
}

func Example() {
	rand.Seed(1338)
	PrintCalculation = false
	n := CreatePerceptron(2, 3, 1)

	dataset := Dataset{
		{Vector{0.0, 0.0}, Vector{1.0}},
		{Vector{1.0, 0.0}, Vector{0.0}},
		{Vector{0.0, 1.0}, Vector{0.0}},
		{Vector{1.0, 1.0}, Vector{1.0}},
	}

	trainer := PerceptronTrainer{&n, dataset}
	trainer.BackPropagation(10000)

	PrintCalculation = true

	n.Calculate(Vector{0.0, 0.0})
	n.Calculate(Vector{1.0, 0.0})
	n.Calculate(Vector{0.0, 1.0})
	n.Calculate(Vector{1.0, 1.0})

	// Output:
	// Input: [0 0] Output: [0.9816677167418877]
	// Input: [1 0] Output: [0.020765305091063144]
	// Input: [0 1] Output: [0.01825325088702373]
	// Input: [1 1] Output: [0.9847884089930483]
}
