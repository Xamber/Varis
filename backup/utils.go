package main

func repeat(f func(), times int) {
	for times > 0 {
		f()
		times--
	}
}

func sum(data []float64) float64 {
	var result float64
	for _, i := range data {
		result += i
	}
	return result
}
