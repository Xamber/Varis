<p align="center">
    <img src="https://user-images.githubusercontent.com/1732107/33759520-a9eb5208-dc13-11e7-9ba2-9c9f97e45ac4.jpg" height="200" alt="Gopher from internet =)" title="Gopher from internet =" />
</p>

# Varis
Neural Networks with GO

[![Build Status](https://travis-ci.org/Xamber/Varis.svg?branch=master)](https://travis-ci.org/Xamber/Varis)
[![Go Report Card](https://goreportcard.com/badge/github.com/Xamber/Varis)](https://goreportcard.com/report/github.com/Xamber/Varis)
[![API Reference](https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667)](https://godoc.org/github.com/Xamber/Varis)
[![codecov](https://codecov.io/gh/Xamber/Varis/branch/master/graph/badge.svg)](https://codecov.io/gh/Xamber/Varis)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/xamber/Varis/blob/master/LICENSE.md)
[![Release](https://img.shields.io/github/tag/xamber/varis.svg?label=latest)](https://github.com/Xamber/Varis/releases/tag/release-0.1)

## About Package
Some time ago I decided to learn Go language and neural networks.
So it's my variation of Neural Networks library. I tried to make library for programmers (not for mathematics).

For now Varis is 0.1 version.

I would be happy if someone can find errors and give advices.
Thank you. Artem.

## Main features
- All neurons and synapses are goroutines.
- Golang channels for connecting neurons.
- No dependencies

## Installation
    go get github.com/Xamber/Varis

## Usage
```go
package main

import (
	"github.com/Xamber/Varis"
)

func main() {
	net := varis.CreatePerceptron(2, 3, 1)

	dataset := varis.Dataset{
		{varis.Vector{0.0, 0.0}, varis.Vector{1.0}},
		{varis.Vector{1.0, 0.0}, varis.Vector{0.0}},
		{varis.Vector{0.0, 1.0}, varis.Vector{0.0}},
		{varis.Vector{1.0, 1.0}, varis.Vector{1.0}},
	}

	trainer := varis.PerceptronTrainer{
		Network: &net,
		Dataset: dataset,
	}

	trainer.BackPropagation(10000)
	varis.PrintCalculation = true

	net.Calculate(varis.Vector{0.0, 0.0}) // Output: [0.9816677167418877]
	net.Calculate(varis.Vector{1.0, 0.0}) // Output: [0.02076530509106318]
	net.Calculate(varis.Vector{0.0, 1.0}) // Output: [0.018253250887023762]
	net.Calculate(varis.Vector{1.0, 1.0}) // Output: [0.9847884089930481]
}

```
## Roadmap 0.2-0.5
- Add locks
- Add training channels
- Improve speed
- Add error return to functions.
- Create more tests and benchmarks.
- Create server and cli realization for use Varis as a application

## Alternatives
[gonn](https://github.com/fxsjy/gonn) | [go-mind](https://github.com/stevenmiller888/go-mind) | [go-perceptron-go](https://github.com/made2591/go-perceptron-go)
