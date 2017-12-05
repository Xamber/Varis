package main

import (
	"github.com/xamber/Varis"
)

func main() {
	n := varis.CreateNetwork(2, 3, 1)

	dataset := varis.Dataset{
		{varis.Vector{0.0, 0.0}, varis.Vector{1.0}},
		{varis.Vector{1.0, 0.0}, varis.Vector{0.0}},
		{varis.Vector{0.0, 1.0}, varis.Vector{0.0}},
		{varis.Vector{1.0, 1.0}, varis.Vector{1.0}},
	}

	varis.BackPropagation(&n, dataset, 10000)
	varis.PrintCalculation = true

	n.Calculate(varis.Vector{0.0, 0.0})
	n.Calculate(varis.Vector{1.0, 0.0})
	n.Calculate(varis.Vector{0.0, 1.0})
	n.Calculate(varis.Vector{1.0, 1.0})
}
