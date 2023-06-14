package main

import (
	"fmt"
	"strings"
)

func contarLetrasPalavras(frase string) map[string]map[rune]int {
	resultado := make(map[string]map[rune]int)

	palavras := strings.Fields(frase)

	for _, palavra := range palavras {
		letras := make(map[rune]int)

		for _, letra := range palavra {
			letras[letra]++
		}

		resultado[palavra] = letras
	}

	return resultado
}

func main() {
	frase := "Olá, este é um exemplo de frase."

	resultado := contarLetrasPalavras(frase)
	fmt.Println(resultado)
}
