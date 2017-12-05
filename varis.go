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
		go func(index int) {
			vector[index] = <-c
			wg.Done()
		}(i)
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

// CreateNetwork make new NN with count of neurons in each Layer.
func CreateNetwork(layers ...int) Network {

	network := Network{}
	network.output = make([]chan float64, 0)

	for index, neurons := range layers {
		layer := []*Neuron{}
		for i := 0; i < neurons; i++ {
			// Standart neuron implimentation
			// callFunc  is neuron.connection.broadcast
			// We will overwrite callbackFunc later
			var neuron = &Neuron{weight: rand.Float64()}
			neuron.callbackFunc = neuron.conn.broadcastSignals
			neuron.collectFunc = neuron.conn.collectSignals

			switch index {
			case 0:
				// Input layer
				// Standart neuron implimentation without callFunc
				neuron.callbackFunc = nil
			case len(layers) - 1:
				// output layer
				// Need to create output channel to redirect Neuron output to NetworkOutput
				outputChan := make(chan float64)
				network.output = append(network.output, outputChan)
				redirect := func(value float64) {
					outputChan <- value
				}
				neuron.callbackFunc = redirect
			}
			layer = append(layer, neuron)
		}
		network.layers = append(network.layers, layer)
	}

	network.ConnectLayers()
	network.RunNeurons()

	return network
}
