package main

import (
	"fmt"
	"time"
)

func main() {
	canal := toWrite("olá mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func toWrite(text string) <-chan string {  // abstração da função main
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
