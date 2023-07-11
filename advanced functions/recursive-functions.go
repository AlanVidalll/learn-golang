package main

import "fmt"

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {
	fmt.Println("Funções recursivas")

	position := uint(10) // neste caso convertemos para uint porque por padrão na inferencia de tipo ele entende como int

	for i := uint(1); i < position; i++ {
		fmt.Println(fibonacci(i))
	}

	fmt.Println(fibonacci(position))

	// 1 1 2 3 5 8 13

	// FUNÇÕES RECURSIVAS NÃO SÃO RECOMENDADAS SE HOUVER MUITAS EXECUÇÕES
}
