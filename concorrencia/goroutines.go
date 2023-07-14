package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARELELISMO
   go toWrite("Olá mundo") // goroutine
      toWrite("Messi e melhor que cr7")  //sem go routine essa função nunca e executada, pois a função acima não finaliza, mas com goroutine ele segue executando o código sem a função acima finalizar
}

func toWrite(text string) {

	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
