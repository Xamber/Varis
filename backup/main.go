package main

import (
	"fmt"
	"math/rand"
)


func main() {

	rand.Seed(1338)

	n := CreateNetwork(2, 2, 2)

	n.ShowStatistic()

	fmt.Println(n.Calculate([]float64{0.0, 0.0}))
	fmt.Println(n.Calculate([]float64{1.0, 1.0}))
	fmt.Println(n.Calculate([]float64{1.0, 0.0}))


}
