package main

import "fmt"

type userType struct {
	name       string
	age        uint8
	adressUser adress // struct aninhada, declarada abaixo
}

type adress struct {
	logradouro string
	number     uint8
}

func main() {

	// maneira explicita de tipagem
	var user userType
	user.name = "alan"
	user.age = 18
	fmt.Println(user)

	// maneira com inferencia de tipo com struct chamando outra struct
	adressExample := adress{"Rua brasil", 45}
	user2 := userType{"alan", 21, adressExample}
	fmt.Println(user2)

	// maneira com inferencia de tipo porem passando somente um valor
	user3 := userType{name: "alan"}
	fmt.Println(user3)
}
