package main // Definição de arquivo principal da aplicação

import (
	"fmt"
	"modulo/auxiliar"
) // Importando packages

// Sintax de função, função main é a função principál da aplicação
func main() {
	fmt.Println("Hello world") // printando hello word
	auxiliar.Escrever()
}

// Para executar a aplicação podemos rodar o comando no terminal => go run <nome do arquivo>


// Para definirmos se uma variavel,função, struct entre outras coisas serão publicaso nomde deve ser com a primeira letra maiuscula e privado será minuscula
// Exemplo func Teste(){} púbiico  func teste(){} privado, para ser importada para outros pacotes deverá ser público