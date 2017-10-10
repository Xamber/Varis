package main

import (
	"fmt"
	"math/rand"
)

func main() {

	rand.Seed(1338)

	n := CreateNetwork(2, 2, 1)
	n.ShowStatistic()

	fmt.Println("Program started..")

}
