package main

import "fmt"

func mesclarMapas(mapa1, mapa2 map[string]int) map[string]int {
	resultado := make(map[string]int)

	for chave, valor := range mapa1 {
		resultado[chave] = valor
	}

	for chave, valor := range mapa2 {
		resultado[chave] = valor
	}

	return resultado

}

func main() {
	mapa1 := map[string]int{
		"maçã":   5,
		"banana": 3,
	}
	mapa2 := map[string]int{
		"maçã":    2,
		"laranja": 4,
		"abacaxi": 1,
	}

	resultado := mesclarMapas(mapa1, mapa2)
	fmt.Println(resultado)
}
