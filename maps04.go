package main

import (
	"fmt"
	"sort"
	"strings"
)

func agruparAnagramas(palavras []string) map[string][]string {
	resultado := make(map[string][]string)

	for _, palavra := range palavras {

		chave := sortString(palavra)

		resultado[chave] = append(resultado[chave], palavra)
	}

	return resultado
}

func sortString(s string) string {
	letters := strings.Split(s, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

func main() {
	palavras := []string{"amor", "roma", "casa", "saco", "mar", "ramo"}

	resultado := agruparAnagramas(palavras)
	fmt.Println(resultado)
}
