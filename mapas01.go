package main

import (
	"fmt"
	"strings"
)

func contarOcorrencias(texto string) map[string]int {

	texto = strings.ToLower(texto)
	texto = strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			return r
		}
		return ' '
	}, texto)

	palavras := strings.Fields(texto)

	ocorrencias := make(map[string]int)

	for _, palavra := range palavras {
		ocorrencias[palavra]++
	}

	return ocorrencias
}

func main() {
	texto := "Olá, este é um exemplo de texto. Olá novamente!"
	resultado := contarOcorrencias(texto)
	fmt.Println(resultado)
}

