package varis

// Function for sum all elements in slice.
func sum(data []float64) float64 {
	var result float64
	for _, i := range data {
		result += i
	}
	return result
}
