package main

import "fmt"

func main() {

	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)
	
//Quanto mais números de vezes vc executar a função worker irá ser mais rapido, somente fique atento ao poder da maquina pois isso limita e vc pode estar chamando funções sem necessidade.
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	

	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}

func worker(tarefas <-chan int, resultados chan<- int) {

	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}
