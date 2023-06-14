package main

import "fmt"

func somarValoresMapa(m map[string]int) int {
	soma := 0

	for _, valor := range m {
		soma += valor
	}

	return soma
}

func main() {
	mapa := map[string]int{
		"maçã":    5,
		"banana":  3,
		"laranja": 4,
	}

	resultado := somarValoresMapa(mapa)
	fmt.Println(resultado)
}
