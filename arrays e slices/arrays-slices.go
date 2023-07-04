package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Arrays e slices so podem receber um tipo de dado, por exemplo não pode receber um int e uma string no mesmo array ou slice

	// declarando array e atribuindo valor de forma explicita
	var array1 [5]string // array sempre tem que declarar o tamanho, caso contrario será um slice
	array1[0] = "Posição 1"
	fmt.Println(array1)

	// declarando array e atribuindo valor usando inferencia de tipo
	array2 := [5]string{"posição 1", "posição 2", "posição 3", "posição 4", "posição 5"}
	fmt.Println(array2)

	array3 := [...]string{"posição 1", "posição 2", "posição 3", "posição 4", "posição 5"} // os 3 pontos seta o valor de tamanho do array pegando o length dele, não é muito utilizado
	fmt.Println(array3)

	// OBS: Array não e muito usado, se faz mais uso de slice.

	slice := []int{1, 2, 3, 4} // quando não passamos nada dentro dos '[]' ele se torna um slice e tamanho dele se torna dinâmico, por isso e muito usado em go e muito mais flexivel
	fmt.Println(slice)
	slice = append(slice, 18)// add um item no slice existente e retorna um novo slice com o valor adicionado
	fmt.Println(slice)

	// reflect.TypeOf verifica o tipo de dado
	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(slice))

	slice2 := array2[1:4]// ele add os valores do indice 1 ao 3 do array
	fmt.Println(slice2)
	
	array2[1] = "posição alterada"
	fmt.Println(slice2) // como podemos ver o valor foi alterado atribuindo um novo valor no array2 e nao no slice, ele funciona como ponteiro

	// ARRAYS INTERNOS

	slice3 := make([]float32,10,11)// refrencia de uma array interno, cria um array de 10 posições coma a capacidade maxima de 15 internamente e retorna uma slice
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // função len e similiar ao length, retorna o tamanho de um slice
	fmt.Println(cap(slice3)) // função cap, retorna a capacidade de um slice

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6) // quando o go indentica que estorou a capacidade maxima estipulada, automaticamente ele cria outro array interno com a capacidade dobrada e retorna um novo slice
	fmt.Println(len(slice3)) // função len e similiar ao length, retorna o tamanho de um slice
	fmt.Println(cap(slice3)) // função cap, retorna a capacidade de um slice

	slice4 := make([]float32,5)// quando n'ao passamos a capacidade ele pega o valor do tamnho de posicoes e seta a capacidade
	fmt.Println(slice4)
	fmt.Println(len(slice4)) 
	fmt.Println(cap(slice4))

	slice4 = append(slice4,6) // quando da estouro ele pega o tamho atual e dobra acapacidade do array
	fmt.Println(len(slice4)) 
	fmt.Println(cap(slice4))


	// RESUMINDO ARRAY E UMA LISTA COM TAMANHO FIXO E UM SLICE COM TAMANHO DINAMICO

}
