package main

import "fmt"

func main() {
	fmt.Println("ponteiros")

	var variable1 int = 10
	var variable2 int = variable1

	fmt.Println(variable1, variable2)
	variable1++
	fmt.Println(variable1, variable2)

	// PONTEIRO
	// Ponteiro é uma referência de memória

	var variable3 int = 10
	var pointer *int = &variable3    // não podemos atribuir um vaor de variável a um ponteiro, para isso temos que passar o '&' na frente da variável
	fmt.Println(variable3, *pointer) // para imprimir o valor de pointer precisamos passar o asteristico, pois sem ele será printado o endereço de memória e não o valor que está dentro desse endereço de memória.

	variable3 = 145
	fmt.Println(variable3, *pointer)
}
