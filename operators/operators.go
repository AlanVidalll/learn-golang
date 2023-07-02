package main

import "fmt"

func main() {

	// ARITMÉTICOS
	fmt.Println(`---ARITMÉTICOS---`)
	sum := 1 + 2
	subtraction := 10 - 5
	division := 10 / 2
	multiplication := 5 * 5
	restDivision := 10 % 3

	fmt.Println(sum, subtraction, division, multiplication, restDivision)
	fmt.Println(`----------------------------------`)

	// ATRIBUIÇÃO
	fmt.Println(`---ATRIBUIÇÃO---`)
	var variavel1 string = "string"
	variavel2 := "string2"
	fmt.Println(variavel1, variavel2)
	fmt.Println(`----------------------------------`)

	// OPERADORES RELACIONAIS
	fmt.Println(`---OPERADORES RELACIONAIS---`)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println(`----------------------------------`)

	// OPERADORES LOGICOS
	fmt.Println(`---OPERADORES LÓGICOS---`)
	fmt.Println(true && false) // operador `and`
	fmt.Println(true || false) // operador `or`
	fmt.Println(!true)         // operador `or`
	fmt.Println(`----------------------------------`)

	// OPERADORES UNÁRIOS
	fmt.Println(`---OPERADORES UNÁRIOS---`)
	number := 10
	number++     // soma de 1 em 1
	number += 10 // soma de 10 em 10
	fmt.Println(number)

	number--    // subtrai de 1 em 1
	number -= 9 // subtrai de 9 em 9
	fmt.Println(number)

	number *= 2 // multiplica por 2
	fmt.Println(number)
	number /= 2 // divide por 2
	fmt.Println(number)
	number %= 2 // divide por 2 e retorna o resto da divisão
	fmt.Println(number)

	fmt.Println(`----------------------------------`)

	fmt.Println(`---OPERADOR TERNÁRIO---`)
	// em go não existe operador ternario
	var text string
	if number > 5 {
		text = "Maior que 5"
	} else {
		text = "Menor que 5"
	}
	fmt.Println(text)
	fmt.Println(`----------------------------------`)

}
