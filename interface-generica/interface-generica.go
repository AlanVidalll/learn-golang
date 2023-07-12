package main

import "fmt"

func generic(interf interface{}) { // Pra quem conhece typescript. isso e bem similiar ao tipo any podendo receber parametro de qualquer tipo
	fmt.Println(interf)
}

func main() {
	fmt.Println("Interface tipo generico")

	generic("teste string")
	generic(58)

	mapa := map[interface{}]interface{}{ // como podemos ver uma verdadeira bagunça recebendo qualquer coisa no map, uma má pratica no geral
		1:            "string",
		"string":     45,
		float32(100): true,
	}

	fmt.Println(mapa)
}
