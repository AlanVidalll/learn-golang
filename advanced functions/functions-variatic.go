package main

import "fmt"

func sum(numbers ...int) int { // função pode receber de varios ints de padrão definido, pode-se receber any parâmetros.
	fmt.Println(numbers) // retorna um slice
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// o conjunto de parâmetros dinâmicos sempre tem q estar por ultimos os fixos sempre vem na frente na declaração
func toWrite(text string, numbers ...int) { // recebe uma string e um conjunto de números com quantidade não definida

	for _, number := range numbers {

		fmt.Println(text, "-", number)

	}
}

func main() {
	fmt.Println("Funções variaticas")

	totalSum := sum(4, 5, 7, 9)
	fmt.Println(totalSum)

	toWrite("Numero", 4, 5, 7, 8, 9, 7, 25)
}
