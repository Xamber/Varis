package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(1338)

	n := CreateNetwork(1, 2, 3, 2, 1)
	n.ShowStatistic()

	n.Calculate( []float64{0.0} )

	fmt.Println("Program started..")
	time.Sleep(time.Second*5)
}
