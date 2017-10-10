package main

import "math"

// Функция вычисляет среднее от среза
func mean(data []float64) float64 {
	total := 0.0
	for _, value := range data {
		total += value
	}

	return total / float64(len(data))
}

// Функция активации (Сигмойда)
func activation_sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// Производная от функции активации
func derivative_sigmoid(x float64) float64 {
	var fx = activation_sigmoid(x)
	return fx * (1 - fx)
}
