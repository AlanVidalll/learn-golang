package main

import "fmt"

func main() {
	fmt.Println("Funções anonimas")

	returnFunction := func(text string) string {

		return fmt.Sprintf("Recebido -> %s", text)
	}("Hello world")

	fmt.Println(returnFunction)
}
