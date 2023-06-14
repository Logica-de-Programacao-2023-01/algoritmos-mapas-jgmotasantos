package main

import (
	"fmt"
)

func somarContagensMapas(mapas []map[string]int) map[string]int {
	resultado := make(map[string]int)

	for _, m := range mapas {
		for palavra, contagem := range m {
			resultado[palavra] += contagem
		}
	}

	return resultado
}

func main() {
	mapas := []map[string]int{
		{"maçã": 2, "banana": 3, "laranja": 1},
		{"maçã": 4, "uva": 2, "laranja": 2},
		{"maçã": 1, "abacaxi": 5, "laranja": 3},
	}

	resultado := somarContagensMapas(mapas)
	fmt.Println(resultado)
}
