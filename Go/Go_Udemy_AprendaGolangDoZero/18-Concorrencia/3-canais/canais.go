package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go Escrever("Olá Mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada!")

	// for {
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	//identifica que o canal fechou para poder passar para o fim do programa.
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func Escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	//fechar o canal para evitar deadlock
	close(canal)
}
