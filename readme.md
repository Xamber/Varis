# Varis
Neural Networks with GO

## Disclaimer
Some time ago I decided to learn Go language and neural networks.
And the best way is to do it - practice. So I use this repository to better understand the language.
I've already encountered many pitfalls in GO, I've learned in practice many tools, data types and "chips" of the language.
I would be happy if someone can find errors and give advice.

Thank you. Artem.

## About package
- No dependencies
- All neurons and synapses are goroutines.
- Golang channels for connecting neurons.
- Neuron have callback function for redirect signal to output or any place you want.

## Installation
    go get https://github.com/Xamber/Varis

## Usage
```go
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

	n.Calculate(varis.Vector{0.0, 0.0}) // [0.9817305250782501]
	n.Calculate(varis.Vector{1.0, 0.0}) // [0.019886950176535658]
	n.Calculate(varis.Vector{0.0, 1.0}) // [0.01911721247536342]
	n.Calculate(varis.Vector{1.0, 1.0}) // [0.9847854383762864]
}
```
## Roadmap
- Improve project structure.
- Add error return to functions.
- Create more tests and benchmarks.
- Create normal documentation.
- Implement Models (Normalize data from bool, integer, array etc.)
- Create signal and weight types.


