package main

import "fmt"

func main() {
	// variável result sendo atribuida com valor de retorno da função add()
	result := add(5, 9)
	fmt.Println(result)

	// declarando uma variaável do tipo função com retorno tipo int
	var f = func(txt string) int {
		fmt.Println(txt)
		return add(6, 9)
	}

	f("funcionou")
	var result2 int = f("funcionou 2")
	fmt.Println(result2)

	// como a função retorna dois valores precisamos declarar suas variáveis para artibuir esses retornos
	resultSum, resultSubtraction := calculations(10, 5)
	fmt.Println(resultSum, resultSubtraction)

	// para pegar um retorno especifico, declaramos a variavel para pegar aquele retorno , e podemos iguinorar o outro passando somente um underline =>  '_'.
	resultSum1, _ := calculations(10, 5)
	fmt.Println(resultSum1)
	// pegando o segundo retorno somente
	_, resultSubtraction1 := calculations(10, 5)
	fmt.Println(resultSubtraction1)
}

// função de somar 2 valores
func add(number1 int, number2 int) int {
	return number1 + number2
}

// No go podemos usar um recurso bem diferente de outras linguagens que podemos ter mais de um retorno na função.
func calculations(number1, number2 int) (int, int) {

	sum := number1 + number2
	subtraction := number1 - number2
	return sum, subtraction

}
