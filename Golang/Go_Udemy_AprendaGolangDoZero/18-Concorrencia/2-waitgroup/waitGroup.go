package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Criando waitGroup
	var waitGroup sync.WaitGroup

	//Adicionando Goroutines (Só informando as quantidades)
	waitGroup.Add(4)

	//função anônima chamada de forma concorrente.
	go func() {
		Escrever("Olá Mundo!")
		//Ao finalizar a função acima, utilza-se o Done() para informar que acabou e retirar um wg da lista
		waitGroup.Done() // -1 do contador...
	}()
	//função anônima chamada de forma concorrente.
	go func() {
		Escrever("Programando em GO!")
		waitGroup.Done() // -1 do contador...
	}()

	go func() {
		Escrever("Goroutine 3")
		waitGroup.Done() // -1
	}()

	go func() {
		Escrever("Goroutine 4")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() //esperar a contagem das goroutines ficarem em zero.

}

func Escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
