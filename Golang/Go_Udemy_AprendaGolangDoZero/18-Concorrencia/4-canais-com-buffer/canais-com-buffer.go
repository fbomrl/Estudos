package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"
	//como só tem 2 valores, caso cinlua um terceiro valor, irá dar deadlock, já que extrapola a capacidade
	// canal <- "Teceiro valor"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
