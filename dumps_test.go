package varis

import (
	"fmt"
	"testing"
)

func TestDumpToJSON(t *testing.T) {
	n := CreateNetwork(2, 3, 1)

	dataset := Dataset{
		{[]float64{0.0, 0.0}, []float64{1.0}},
		{[]float64{1.0, 0.0}, []float64{0.0}},
		{[]float64{0.0, 1.0}, []float64{0.0}},
		{[]float64{1.0, 1.0}, []float64{1.0}},
	}

	BackPropagation(&n, dataset, 10000)

	fmt.Println("After training")
	fmt.Println(0.0, 0.0, "-", n.Calculate(0.0, 0.0))
	fmt.Println(1.0, 0.0, "-", n.Calculate(1.0, 0.0))
	fmt.Println(0.0, 1.0, "-", n.Calculate(0.0, 1.0))
	fmt.Println(1.0, 1.0, "-", n.Calculate(1.0, 1.0))

	js := ToJSON(n)

	n2 := FromJSON(js)
	fmt.Println(0.0, 0.0, "-", n2.Calculate(0.0, 0.0))
	fmt.Println(1.0, 0.0, "-", n2.Calculate(1.0, 0.0))
	fmt.Println(0.0, 1.0, "-", n2.Calculate(0.0, 1.0))
	fmt.Println(1.0, 1.0, "-", n2.Calculate(1.0, 1.0))

}
