package varis

import (
	"fmt"
	"testing"
)

func BenchmarkNetwork(b *testing.B) {

	for i := 0; i < b.N; i++ {
		n := CreateNetwork(2, 3, 1)

		dataset := Dataset{
			{[]float64{0.0, 0.0}, []float64{1.0}},
			{[]float64{1.0, 0.0}, []float64{0.0}},
			{[]float64{0.0, 1.0}, []float64{0.0}},
			{[]float64{1.0, 1.0}, []float64{1.0}},
		}

		trainer := Trainer{&n, BackPropagation}
		trainer.TrainByDataset(dataset, 10000)

		fmt.Println("After training")
		fmt.Println(0.0, 0.0, "-", n.Calculate(0.0, 0.0))
		fmt.Println(1.0, 0.0, "-", n.Calculate(1.0, 0.0))
		fmt.Println(0.0, 1.0, "-", n.Calculate(0.0, 1.0))
		fmt.Println(1.0, 1.0, "-", n.Calculate(1.0, 1.0))
	}
}
