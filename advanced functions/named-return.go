package main

import "fmt"

func calculated(n1, n2 int) (sum int, subtraction int) {
	sum = n1 + n2
	subtraction = n1 - n2

	return
}

func main() {
	fmt.Println("Funções nomeadas")
	
	sum, subtraction := calculated(10, 5)
	fmt.Println(sum, subtraction)
}
