package main

import "fmt"

func torecover() {
	// função de recuperação
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}

}

func approvedStudent(n1, n2 float32) bool {
	defer torecover() // para chamar a função de recuperar temos que usar o defer
	average := (n1 + n2) / 2

	if average > 6 {

		return true
	} else if average < 6 {
		return false
	}

	panic("A MÉDIA É EXTAMENTE 6!") // panic irá parar tudo execução mas com recover() podemos recuperar  a aplicação
}

func main() {

	fmt.Println(approvedStudent(6, 6))
	fmt.Println("Pós execução!")

}
