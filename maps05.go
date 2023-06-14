package main

import (
	"fmt"
)

func contarCaracteres(texto string) map[string]int {
	resultado := make(map[string]int)

	for _, caractere := range texto {
		chave := string(caractere)
		resultado[chave]++
	}

	return resultado
}

func main() {
	texto := "abacaxi"

	resultado := contarCaracteres(texto)
	fmt.Println(resultado)
}
