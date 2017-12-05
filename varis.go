package varis

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var PrintCalculation = false

type Vector []float64
type neuronFunction func(x float64) float64

func (v Vector) sum() (result float64) {
	for _, i := range v {
		result += i
	}
	return
}

func (v Vector) Broadcast(channels []chan float64) {
	if len(v) != len(channels) {
		panic("Lenght not equal")
	}

	for i, c := range channels {
		c <- v[i]
	}
}

func CollectVector(channels []chan float64) (vector Vector) {
	count := len(channels)
	vector = make(Vector, count)

	wg := sync.WaitGroup{}
	wg.Add(count)

	for i, c := range channels {
		go func(index int, channel chan float64) {
			vector[index] = <-channel
			wg.Done()
		}(i, c)
	}

	wg.Wait()
	return vector
}

// ACTIVATION store default activation function.
var ACTIVATION neuronFunction = func(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// DEACTIVATION store default deactivation function.
var DEACTIVATION neuronFunction = func(x float64) float64 {
	var fx = ACTIVATION(x)
	return fx * (1 - fx)
}
