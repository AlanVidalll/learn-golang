package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 2 {
		i++
		fmt.Println("Incrementando i", i)
		time.Sleep(time.Second)
	}

	fmt.Println("Incrementando i", i) // nessa primeira condição de for temos a variavel i acessível fora dele

	for j := 0; j < 2; j++ {

		fmt.Println("Incrementando j", j) // nessa estrutura de for não conseguimos ter acessoa variável j fora do escopo do for
		time.Sleep(time.Second)
	}

	names := [3]string{"Alan", "João", "Paulo"}

	for index, name := range names { // for range usado para interar com arrays, strings , maps
		fmt.Println(index, "-", name)
	}

	for index, letter := range "Palavra" { // para interar com caracteres por padrão ele retorna os números da tabela ASCII
		fmt.Println(index, "-", string(letter)) // para converter o numero da tabela ASCII precimos usar a função string()
	}

	users := map[string]string{"name": "Alan", "cor": "branco", "sexo": "masculino"}

	for key, value := range users { //interando com map
		fmt.Println(key, "-", value)
	}

	//NÃO E POSSÍVEL INTERAR COM UMA STRUCT
}
