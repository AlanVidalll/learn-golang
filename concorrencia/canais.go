package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARELELISMO
	
	canal := make(chan string)
	go toWrite("Olá mundo", canal)

	/*for {
		mensagem, aberto := <-canal // recebendo um valor para o codigo seguir abaixo
		if !aberto {                // verifica se o canal esta aberto
			break
		}
		fmt.Println(mensagem)
	}*/  // função e a mesma que debaixo 

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}

func toWrite(text string, canal chan string) {

	for i := 0; i < 5; i++ {
		canal <- text // mandando um valor para dentro do canal
		time.Sleep(time.Second)
	}

	close(canal) // fecha o canal
}
