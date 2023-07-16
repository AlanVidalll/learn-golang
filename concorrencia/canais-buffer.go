package main

import "fmt"

func main() {
	fmt.Println("Canal com buffer")

	channel := make(chan string, 2) // define a capacidade do canal , se adicionarmos um terceiro intem irá dar deadlock, 

	channel <- "Olá mundo"
	channel <- "Programando em go"

	mensagem := <-channel
	mensagem2 := <-channel

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
