package varis

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	vector := Vector{1.1, 2.2, 3.3}
	summa := vector.sum()
	if summa != 6.6 {
		t.Error("Sums not equal")
	}
}

func TestZeroSum(t *testing.T) {
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

func TestCreation(t *testing.T) {
	_ = CreateNetwork(2, 3, 1)
}

func TestDumpToJSON(t *testing.T) {
	n1 := CreateNetwork(2, 3, 1)

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
