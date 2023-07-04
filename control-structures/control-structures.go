package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	number := 10

	if number > 15 { // estrutura basica de if e else
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("menor ou igual a 15")
	}

	if otherNumber := number; otherNumber > 0 { // if init, inicia uma variavel no if e verifica a condição da mesma
		fmt.Println("Maior que 0") // vantagem de criar a variavel aqui, não conseguimos acessar ela fora do escopo do if
	} else if otherNumber < -10 {
		fmt.Println("Numero menor que 0 ")
	} else {
		fmt.Println("entre 0 e -10")
	}
}
