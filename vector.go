package varis

import (
	"sync"
)

type Vector []float64

func (v Vector) sum() (result float64) {
	for _, i := range v {
		result += i
	}
	return result
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
