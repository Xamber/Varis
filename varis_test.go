package varis

import (
	"math/rand"
	"testing"
)

func TestVectorSum(t *testing.T) {
	vector := Vector{1.1, 2.2, 3.3}
	summa := vector.sum()
	if summa != 6.6 {
		t.Error("Sums not equal")
	}
}

func TestVectorZeroSum(t *testing.T) {
	vector := Vector{}
	summa := vector.sum()
	if summa != 0.0 {
		t.Error("Sums not equal")
	}
}

func TestGenerateUUID(t *testing.T) {
	uuid := generate_uuid()
	if len(uuid) != 36 {
		t.Error("Len of uuid not a same:", len(uuid))
	}
}

func TestNetworkFunction(t *testing.T) {
	_ = CreatePerceptron(2, 3, 1)
}

func TestDeactivationFunction(t *testing.T) {
	if DEACTIVATION(0.123123123) != 0.24905493220237876 {
		t.Error("Wrong deactivation function")
	}
}

func TestActivationFunction(t *testing.T) {
	if ACTIVATION(0.123123123) != 0.5307419550064931 {
		t.Error("Wrong activation function")
	}
}

func TestNetworkTrain(t *testing.T) {
	nn := CreatePerceptron(2, 3, 1)

	dataset := Dataset{
		{Vector{0.0, 0.0}, Vector{1.0}},
		{Vector{1.0, 0.0}, Vector{0.0}},
		{Vector{0.0, 1.0}, Vector{0.0}},
		{Vector{1.0, 1.0}, Vector{1.0}},
	}

	BackPropagation(&nn, dataset, 10000)

	if nn.Calculate(Vector{0.0, 0.0})[0] < 0.800000 {
		t.Log("Wrong with calculation: ", Vector{0.0, 0.0}, " result: ", nn.Calculate(Vector{0.0, 0.0})[0])
	}

	if nn.Calculate(Vector{1.0, 0.0})[0] > 0.200000 {
		t.Log("Wrong with calculation: ", Vector{1.0, 0.0}, " result: ", nn.Calculate(Vector{1.0, 0.0})[0])
	}

	if nn.Calculate(Vector{0.0, 1.0})[0] > 0.200000 {
		t.Log("Wrong with calculation: ", Vector{0.0, 1.0}, " result: ", nn.Calculate(Vector{0.0, 1.0})[0])
	}

	if nn.Calculate(Vector{1.0, 1.0})[0] < 0.800000 {
		t.Log("Wrong with calculation: ", Vector{1.0, 1.0}, " result: ", nn.Calculate(Vector{1.0, 1.0})[0])
	}

}

func TestNetworkCalculate(t *testing.T) {
	rand.Seed(1488)
	nn := CreatePerceptron(2, 3, 1)

	if nn.Calculate(Vector{0.0, 0.0})[0] != 0.6631222149019123 {
		t.Error("Wrong with calculation")
	}
}

func TestDumpToJSON(t *testing.T) {
	n1 := CreatePerceptron(2, 3, 1)

	dataset := Dataset{
		{Vector{0.0, 0.0}, Vector{1.0}},
		{Vector{1.0, 0.0}, Vector{0.0}},
		{Vector{0.0, 1.0}, Vector{0.0}},
		{Vector{1.0, 1.0}, Vector{1.0}},
	}

	BackPropagation(&n1, dataset, 10000)
	json := ToJSON(n1)
	n2 := FromJSON(json)

	if n1.Calculate(Vector{0.0, 0.0})[0] != n2.Calculate(Vector{0.0, 0.0})[0] {
		t.Error("Result of neural newtwork calculating not a same ")
	}

	if n1.Calculate(Vector{1.0, 0.0})[0] != n2.Calculate(Vector{1.0, 0.0})[0] {
		t.Error("Result of neural newtwork calculating not a same ")
	}

	if n1.Calculate(Vector{0.0, 1.0})[0] != n2.Calculate(Vector{0.0, 1.0})[0] {
		t.Error("Result of neural newtwork calculating not a same ")
	}

	if n1.Calculate(Vector{1.0, 1.0})[0] != n2.Calculate(Vector{1.0, 1.0})[0] {
		t.Error("Result of neural newtwork calculating not a same ")
	}
}
