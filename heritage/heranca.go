package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      uint8
	height   uint8
}

type student struct {
	person  // herança em go so declara a struct que ele atribui os campos de pessoa dentro da struct student
	course  string
	college string
}

func main() {

	person1 := person{"Alan", "Fonseca", 25, 178}
	student1 := student{person1, "Direito", "Unesp"}
	fmt.Println(student1)
	fmt.Println(student1.height) // não precisa acessar student1.person1.height podemos acessar direto student1.height
}
