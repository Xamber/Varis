package varis

import (
	"sync"
)

// Vector implement array of float64
type Vector []float64

// sum all element of vector
func (v Vector) sum() (result float64) {
	for _, i := range v {
		result += i
	}
	return result
}

// broadcast send all values to each of channels of channels array
func (v Vector) broadcast(channels []chan float64) {
	if len(v) != len(channels) {
		panic("Length not equal")
	}

	for i, c := range channels {
		c <- v[i]
	}
}

// collectVector get all values from each of channels of channels array
func collectVector(channels []chan float64) (vector Vector) {
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

// is compate two Vectors
func (v Vector) is(other Vector) bool {
	if len(v) != len(other) {
		return false
	}

	for index, value := range v {
		if value != other[index] {
			return false
		}
	}

	return true
}
