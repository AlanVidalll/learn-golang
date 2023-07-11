package main

import "fmt"

func closure() func() {
	text := "Dentro da função closure"

	function := func() {
		fmt.Println(text)

	}

	return function
}

func main() {
	fmt.Println("Funções clousure")
	text := "Dentro da função main"
	fmt.Println(text)

	newFunction := closure()
	newFunction() // irá imprimir o que foi declarada na função closure e não na main
}
