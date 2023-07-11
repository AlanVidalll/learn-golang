package main

import "fmt"

func function1() {
	fmt.Println("Executando a função 1")
}

func function2() {
	fmt.Println("Executando a função 2")
}

func approvedStudent(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado") // usando defer ele vai executar antes de executar o return evitando a duplicidade do código na linhas comentada
	fmt.Println("Entrando na função para ver se o aluno está aprovado")

	average := (n1 + n2) / 2

	if average >= 7 {
		//fmt.Println("Média calculada. Resultado será retornado") // esse print quermos que mostre independente do resultado, padrão normal sem defer teriamos que chamar duas vezes função
		return true
	}
	//fmt.Println("Média calculada. Resultado será retornado")
	return false
}

func main() {
	defer function1()
	// DEFER = ADIAR
	function2()

	fmt.Println(approvedStudent(7, 8))

}
