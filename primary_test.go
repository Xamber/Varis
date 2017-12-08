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

func TestModel(t *testing.T) {
	n := CreatePerceptron(2, 3, 1)

	trainer := PerceptronTrainer{&n, dataset}
	trainer.BackPropagation(10000)

	PrintCalculation = true

	n.Calculate(Vector{0.0, 0.0})
	n.Calculate(Vector{1.0, 0.0})
	n.Calculate(Vector{0.0, 1.0})
	n.Calculate(Vector{1.0, 1.0})

	type Model struct {
		Network *Perceptron
		X1      BooleanField `direction:"input"`
		X2      BooleanField `direction:"input"`
		O       BooleanField `direction:"output"`
	}

	f := Model{Network: &n}

	calculate := GenerateRunFunction(f)

	calculate([]interface{}{false, false})
	calculate([]interface{}{true, false})
	calculate([]interface{}{false, true})
	calculate([]interface{}{true, true})
}

func Example() {
	rand.Seed(1338)
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

	// Model example section
	type Model struct {
		Network *Perceptron
		X1      BooleanField `direction:"input"`
		X2      BooleanField `direction:"input"`
		O       BooleanField `direction:"output"`
	}

	f := Model{Network: &n}

	calculate := GenerateRunFunction(f)

	calculate([]interface{}{false, false})
	calculate([]interface{}{true, false})
	calculate([]interface{}{false, true})
	calculate([]interface{}{true, true})

	// Output:
	// Input: [0 0] Output: [0.9816677167418877]
	// Input: [1 0] Output: [0.02076530509106318]
	// Input: [0 1] Output: [0.018253250887023762]
	// Input: [1 1] Output: [0.9847884089930481]
	// Input: [0 0] Output: [0.9816677167418877]
	// Input: [1 0] Output: [0.02076530509106318]
	// Input: [0 1] Output: [0.018253250887023762]
	// Input: [1 1] Output: [0.9847884089930481]
}
