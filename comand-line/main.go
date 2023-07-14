package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	aplication := app.Generated()
	if error := aplication.Run(os.Args); error != nil{
		log.Fatal(error)
	}
	
}
