package main

import (
	"fmt"
)

func main() {

	// Maneira explicita de declarar variável e sua tipagem.
	var variable1 string = "Variavel 1"
	fmt.Println(variable1)

	// Maneira abstrata de declarar variável e sua tipragem, conhecido como inferência de tipo, "Atribuição do tipo determinado no seu valor".
	variable2 := "Variavel 2"
	fmt.Println(variable2)

	// podemos declarar multiplas variáveis em bloco conforme esse exemplo e usar normalmente, essa e a maneira explícita.
	var (
		variable3 string = "Variavel 3"
		variable4 int    = 10
		variable5 bool   = false
	)
	fmt.Println(variable3, variable4, variable5)

	// Declarando multiplas variáveis com inferência de tipo, "Atribuição do tipo determinado no seu valor".
	variable6, variable7, variable8 := "Variavel 6", 10, false
	fmt.Println(variable6, variable7, variable8)

	// Todos esees exemplos são aplicavéis para constatnte também.
	const variable9 = "Variavel 9"
	variable10 := "Variavel 10"
	fmt.Println(variable9, variable10)

	// aqui e um exemplo de troca de valores de duas variáveis, em outra libguagens normalmente fariamos com uma variável auxilia, mas no go pode se dispensar essa pratica e fazermos conforme abaixo:
	variable1, variable2 = variable2, variable1
	fmt.Println(variable1, variable2)
}
