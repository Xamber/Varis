package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(1337)

	n := CreateNetwork(2, 4, 1)

	fmt.Println(0.0, 0.0, "-", n.Calculate([]float64{0.0, 0.0}))
	fmt.Println(1.0, 0.0, "-", n.Calculate([]float64{1.0, 0.0}))
	fmt.Println(0.0, 1.0, "-", n.Calculate([]float64{0.0, 1.0}))
	fmt.Println(1.0, 1.0, "-", n.Calculate([]float64{1.0, 1.0}))

	repeat(func() {
		n.Train([]float64{0.0, 0.0}, []float64{1.0})
		n.Train([]float64{1.0, 0.0}, []float64{0.0})
		n.Train([]float64{0.0, 1.0}, []float64{0.0})
		n.Train([]float64{1.0, 1.0}, []float64{1.0})
	}, 10000)

	fmt.Println(0.0, 0.0, "-", n.Calculate([]float64{0.0, 0.0}))
	fmt.Println(1.0, 0.0, "-", n.Calculate([]float64{1.0, 0.0}))
	fmt.Println(0.0, 1.0, "-", n.Calculate([]float64{0.0, 1.0}))
	fmt.Println(1.0, 1.0, "-", n.Calculate([]float64{1.0, 1.0}))
}
