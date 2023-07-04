package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":     "Alan",
		"lastName": "Fonseca",
	}

	fmt.Println(user)
	fmt.Println(user["name"]) // diferentemente do struct, não consiguimos acessar user.name, devemos fazer user["name"]

	user2 := map[string]map[string]string{ // aninhamento de map
		"name": {
			"firstName": "Alan",
			"lastName":  "Fonseca",
		},
		"course": {
			"name":    "Engenharia",
			"college": "Unesp",
		},
	}

	fmt.Println(user2)
	delete(user2, "name") // deleta a chave curso
	fmt.Println(user2)
	user2["name"] = map[string]string{
		"firstName": "Alan",
		"lastName":  "Fonseca",
	}
	fmt.Println(user2)
}
