package main

import "fmt"

var number int

func main() {
	fmt.Println("Função init")
	fmt.Println("Função main sendo executada")
	fmt.Println(number) // como pode-se ver ela foi declarada fora das funções porém foi atribuida o valor na função init e o retorno do print no main e o valor atribuido no init
}

func init() {
	number = 25

	fmt.Println("Função init sendo executada") // essa função é executada antes do main
}

// a função init pode ser decrarada por arquivo e antes de carregar o main todas funções init será executada
