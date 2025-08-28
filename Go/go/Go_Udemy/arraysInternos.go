package main

import "fmt"

func main() {

	//Arrays Internos

	//tipo / tamanho / capacidade
	slice3 := make([]float32, 10, 15)
	fmt.Println("------------------")
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade
	//Quando a capacidade é "estourada"o próprio sistema cria um novo array com o dobro da capacidade, assim evitando que o sistema "uquebre"
	fmt.Println("------------------")
	//tipo / tamanho

	//quando não é definido a capacidade do slice ele define por padrão igual ao tamanho dele.
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) //tamanho
	fmt.Println(cap(slice4)) //capacidade
}
