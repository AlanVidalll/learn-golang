package main

import "fmt"

type user struct {
	name string
	age uint
}

func toWrite() { // função normal ela fica solta no código
	fmt.Println("Escrevendo")
}


func (u user) salvar(){
	fmt.Printf("Sanvaldo usuário %s no banco de dados\n", u.name)
}

func (u user) maiorDeIdade() bool {  // metodos segue os mesmos conceitos de função podendo colocar retorno também
	return u.age >= 18
}

func (u *user) fazerAniversario()  { // Aqui passamos como ponteiro pois o valor fica armazenado podendo ser alterado e chamdo em qualquer momento com valor atualizado
	u.age++  // senão passarmos como ponteiro e chamar o metodo no main por exemplo e printa ele não vai ter alteração de valor
}

func main() {
	fmt.Println("Metodos")

	user1 := user{"alan", 33}
	user1.salvar()

	user2 := user{"alan", 15}
	userReturn := user2.maiorDeIdade()
	fmt.Println(userReturn)

	user2.fazerAniversario() // com ponteiro o valor sera acrecentado e o print abaixo retornara 16 sem ponteiro retornara 15
	fmt.Println(user2.age)
}