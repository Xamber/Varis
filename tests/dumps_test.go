package varis_tests

import (
	"testing"
	"github.com/xamber/Varis"
)

func TestDumpToJSON(t *testing.T) {
	n1 := varis.CreateNetwork(2, 3, 1)

	dataset := varis.Dataset{
		{[]float64{0.0, 0.0}, []float64{1.0}},
		{[]float64{1.0, 0.0}, []float64{0.0}},
		{[]float64{0.0, 1.0}, []float64{0.0}},
		{[]float64{1.0, 1.0}, []float64{1.0}},
	}

	varis.BackPropagation(&n1, dataset, 10000)
	json := varis.ToJSON(n1)
	n2 := varis.FromJSON(json)

	if n1.Calculate(0.0, 0.0)[0] != n2.Calculate(0.0, 0.0)[0] {
		t.Error("Result of neural newtwork calculating not a same ")
	}

	if n1.Calculate(1.0, 0.0)[0] != n2.Calculate(1.0, 0.0)[0] {
		t.Error("Result of neural newtwork calculating not a same ")
	}

	if n1.Calculate(0.0, 1.0)[0] != n2.Calculate(0.0, 1.0)[0] {
		t.Error("Result of neural newtwork calculating not a same ")
	}

	if n1.Calculate(1.0, 1.0)[0] != n2.Calculate(1.0, 1.0)[0] {
		t.Error("Result of neural newtwork calculating not a same ")
	}

}
