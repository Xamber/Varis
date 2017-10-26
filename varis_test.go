package varis

import (
	"fmt"
)

func ExampleNetwork() {
	n := CreateNetwork(2, 3, 1)

	dataset := Dataset{
		{[]float64{0.0, 0.0}, []float64{1.0}},
		{[]float64{1.0, 0.0}, []float64{0.0}},
		{[]float64{0.0, 1.0}, []float64{0.0}},
		{[]float64{1.0, 1.0}, []float64{1.0}},
	}

	trainer := Trainer{&n, BackPropagation}
	trainer.TrainByDataset(dataset, 10000)

	fmt.Println(0.0, 0.0, "-", round(n.Calculate(0.0, 0.0)[0]))
	fmt.Println(1.0, 0.0, "-", round(n.Calculate(1.0, 0.0)[0]))
	fmt.Println(0.0, 1.0, "-", round(n.Calculate(0.0, 1.0)[0]))
	fmt.Println(1.0, 1.0, "-", round(n.Calculate(1.0, 1.0)[0]))
	// Output:
	// 0 0 - 1
	// 1 0 - 0
	// 0 1 - 0
	// 1 1 - 1
}
