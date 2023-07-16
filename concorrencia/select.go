package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {

			time.Sleep(time.Millisecond * 500)
			channel1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Canal 2"

		}
	}()

	for {
		select { // select usado somente com concorrencia, similiar ao switch
		case mensagemCanal1 := <-channel1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-channel2:
			fmt.Println(mensagemCanal2)
		}

	}
}

// Se não usarmos select o canal 1 sempre vai esperar o canal 2 finalizar.
