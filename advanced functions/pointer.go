package main

import "fmt"

func signalInverter(number int) int { // função sem ponteiro

	return number * -1
}

func signalInverterwithPointer(number *int) { // função com ponteiro

	*number = *number * -1

}

func main() {
	fmt.Println("Funções com ponteiros")
	// quando passamos a variavel number para a função ele esta passando uma copia se imprimirmos a variavel o valor dela será 20
	number := 20
	inverterNumber := signalInverter(number)
	fmt.Println(inverterNumber)
	fmt.Println(number) //retorna 20 positivo

	newNumber := 40
	fmt.Println(newNumber)
	signalInverterwithPointer(&newNumber)
	fmt.Println(newNumber) // o valor continua negativo porque não se cria uma copia e sim altera o valor no endereço de memória
}
