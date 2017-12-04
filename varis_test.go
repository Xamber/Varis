package varis

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	elems := []float64{1.1, 2.2, 3.3}
	summa := sum(elems)
	if summa != 6.6 {
		t.Error("Sums not equal")
	}
}

func TestZeroSum(t *testing.T) {
	elems := []float64{}
	summa := sum(elems)
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
		{[]float64{0.0, 0.0}, []float64{1.0}},
		{[]float64{1.0, 0.0}, []float64{0.0}},
		{[]float64{0.0, 1.0}, []float64{0.0}},
		{[]float64{1.0, 1.0}, []float64{1.0}},
	}

	BackPropagation(&n1, dataset, 10000)
	json := ToJSON(n1)
	n2 := FromJSON(json)

	if n1.Calculate([]float64{0.0, 0.0})[0] != n2.Calculate([]float64{0.0, 0.0})[0] {
		t.Error("Result of neural newtwork calculating not a same ")
	}

	if n1.Calculate([]float64{1.0, 0.0})[0] != n2.Calculate([]float64{1.0, 0.0})[0] {
		t.Error("Result of neural newtwork calculating not a same ")
	}

	if n1.Calculate([]float64{0.0, 1.0})[0] != n2.Calculate([]float64{0.0, 1.0})[0] {
		t.Error("Result of neural newtwork calculating not a same ")
	}

	if n1.Calculate([]float64{1.0, 1.0})[0] != n2.Calculate([]float64{1.0, 1.0})[0] {
		t.Error("Result of neural newtwork calculating not a same ")
	}
}

func BenchmarkNetwork(b *testing.B) {

	for i := 0; i < b.N; i++ {
		n := CreateNetwork(2, 3, 1)

		dataset := Dataset{
			{[]float64{0.0, 0.0}, []float64{1.0}},
			{[]float64{1.0, 0.0}, []float64{0.0}},
			{[]float64{0.0, 1.0}, []float64{0.0}},
			{[]float64{1.0, 1.0}, []float64{1.0}},
		}

		BackPropagation(&n, dataset, 10000)

		fmt.Println("After training")
		fmt.Println(0.0, 0.0, "-", n.Calculate([]float64{0.0, 0.0}))
		fmt.Println(1.0, 0.0, "-", n.Calculate([]float64{1.0, 0.0}))
		fmt.Println(0.0, 1.0, "-", n.Calculate([]float64{0.0, 1.0}))
		fmt.Println(1.0, 1.0, "-", n.Calculate([]float64{1.0, 1.0}))
	}
}

func BenchmarkLargeNetwork(b *testing.B) {

	for i := 0; i < b.N; i++ {
		n := CreateNetwork(2, 4, 8, 16, 32, 64, 64, 32, 16, 8, 4, 2, 1)

		dataset := Dataset{
			{[]float64{0.0, 0.0}, []float64{1.0}},
			{[]float64{1.0, 0.0}, []float64{0.0}},
			{[]float64{0.0, 1.0}, []float64{0.0}},
			{[]float64{1.0, 1.0}, []float64{1.0}},
		}

		BackPropagation(&n, dataset, 10)

		fmt.Println("After training")
		fmt.Println(0.0, 0.0, "-", n.Calculate([]float64{0.0, 0.0}))
		fmt.Println(1.0, 0.0, "-", n.Calculate([]float64{1.0, 0.0}))
		fmt.Println(0.0, 1.0, "-", n.Calculate([]float64{0.0, 1.0}))
		fmt.Println(1.0, 1.0, "-", n.Calculate([]float64{1.0, 1.0}))
	}
}
