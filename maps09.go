package main

import (
	"fmt"
)

func fibonacci(n int) map[int]int {
	resultado := make(map[int]int)

	resultado[0] = 0

	if n > 0 {
		resultado[1] = 1
	}

	for i := 2; i <= n; i++ {
		resultado[i] = resultado[i-1] + resultado[i-2]
	}

	return resultado
}

func main() {
	n := 10

	resultado := fibonacci(n)
	fmt.Println(resultado)
}
