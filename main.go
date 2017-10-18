package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(1338)

	n := CreateNetwork(2, 2, 1)

	fmt.Println("Before training")
	fmt.Println(0.0, 0.0, "-", n.Calculate(0.0, 0.0))
	fmt.Println(1.0, 0.0, "-", n.Calculate(1.0, 0.0))
	fmt.Println(0.0, 1.0, "-", n.Calculate(0.0, 1.0))
	fmt.Println(1.0, 1.0, "-", n.Calculate(1.0, 1.0))

	dataset := Dataset{
		Frame{[]float64{0.0, 0.0}, []float64{1.0}},
		Frame{[]float64{1.0, 0.0}, []float64{0.0}},
		Frame{[]float64{0.0, 1.0}, []float64{0.0}},
		Frame{[]float64{1.0, 1.0}, []float64{1.0}},
	}

	trainer := Trainer{&n, BackPropagation}
	trainer.TrainByDataset(dataset, 10000)

	fmt.Println("After training")
	fmt.Println(0.0, 0.0, "-", n.Calculate(0.0, 0.0))
	fmt.Println(1.0, 0.0, "-", n.Calculate(1.0, 0.0))
	fmt.Println(0.0, 1.0, "-", n.Calculate(0.0, 1.0))
	fmt.Println(1.0, 1.0, "-", n.Calculate(1.0, 1.0))
}
