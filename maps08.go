package main

import (
	"fmt"
)

func equalizarGastos(gastos map[string]float64) map[string]float64 {
	total := 0.0

	for _, valor := range gastos {
		total += valor
	}

	media := total / float64(len(gastos))

	resultado := make(map[string]float64)
	for pessoa, valor := range gastos {
		resultado[pessoa] = valor - media
	}

	return resultado
}

func main() {
	gastos := map[string]float64{
		"Alice":   50.0,
		"Bob":     30.0,
		"Charlie": 20.0,
		"Dave":    40.0,
	}

	resultado := equalizarGastos(gastos)
	fmt.Println(resultado)
}
