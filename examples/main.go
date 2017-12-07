package main

import (
	"github.com/xamber/Varis"
)

func main() {
	n := varis.CreatePerceptron(2, 3, 1)

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

	// Model example section
	type Model struct {
		Network *varis.Perceptron
		X1      varis.BooleanField `direction:"input"`
		X2      varis.BooleanField `direction:"input"`
		O       varis.BooleanField `direction:"output"`
	}

	f := Model{Network: &n}

	calculate := varis.GenerateRunFunction(f)

	calculate([]interface{}{false, false})
	calculate([]interface{}{true, false})
	calculate([]interface{}{false, true})
	calculate([]interface{}{true, true})
}
