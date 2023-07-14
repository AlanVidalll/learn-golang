package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARELELISMO

	var waitGroup sync.WaitGroup
	// sincronizando routines

	waitGroup.Add(2)

	go func() {
		toWrite("Olá mundo") // goroutine
		waitGroup.Done()     // -1
	}()

	go func() {
		toWrite("Messi e melhor que cr7") // goroutine
		waitGroup.Done()                  // -1
	}()

	waitGroup.Wait() // 0
}

func toWrite(text string) {

	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
