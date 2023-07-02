package main

import (
	"errors"
	"fmt"
)

func main() {

	// TIPOS INTEIROS
	//int8, int16, int32, int64.

	var number int8 = 100
	fmt.Println(number)

	// podemos declarar da seguinte maneira para que seja atribuida de forma automatica o tipo de dado.
	var number1 int = 100 // ou
	number2 := 200
	fmt.Println(number1, number2)

	// o int ele aceita números negativos, temos um tipo chamado uint que não aceita números negativo e gerando um erro ao declarar um número negativo.
	//var number4 uint = -56 => gera erro e não compila.
	// ele tem a mesma convenção do int tendo 4 tipos => uint8, uint16, uint32 e uint64.
	number4 := 56
	fmt.Println(number4)

	// Alias
	// int32 = rune
	var number5 rune = 1000
	fmt.Println(number5)

	// Alias
	// uint = byte
	var number6 byte = 100
	fmt.Println(number6)

	// números reais existem 2 tipos.
	//float32 e float64
	// aqui no float temos que declarar um dos 2, diferentemente do int não podemos so declarar float para que ele faça uma atribuição automática.
	// podemos usar a inferência de tipo para que ele atribua automaticamente

	var number7 float32 = 100.78
	fmt.Println(number7)

	number8 := 90.50
	fmt.Println(number8)

	// TIPOS STRINGS

	var str string = "string"
	fmt.Println(str)

	// não existe char no go, o char ele sempre retorna um int veja como funciona abaixo:
	char := 'B'
	fmt.Println(char) // ele retorno  o número da tabela ascii no caso é 66

	// TIPO BOOLEANO

	bool := true
	fmt.Println(bool)

	// TIPO DE ERRO

	var error1 error
	fmt.Println(error1)

	// para atribuirmos um valor ao tipo error temos que importar um pacote chamado 'errors' como exemplo abaixo
	// var error1 error = "Erro interno" => não funciona gera um erro
	var error2 error = errors.New("Erro interno")
	fmt.Println(error2)

	// TODOS OS DADOS TEM UM VALOR INICIAL CASO NÃO SE ATRIBUA VALOR NELE SEGUE A LISTA ABAIXO:
	// string retorna string vazia ""
	// number retorna 0
	// bool retorna false
	// error retorna <nil> o que seria similiar ao null de outras linguagens
}
