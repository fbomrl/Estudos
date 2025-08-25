package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")
	//Quando atribui um valor para uma variável, esse valor é uma cópia, por isso o valor de variável2 não aumentou após o ++
	var variavel1 int = 10
	var variavel2 int = variavel1

	variavel1++
	fmt.Println(variavel1, variavel2)

	//Ponteiro é uma referencia de memória
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	//Quando rota aqui mostra o endereço de memória que está no ponteiro e mostra o valor armazenado na variavel3
	fmt.Println(variavel3, ponteiro)

	//desreferenciação, colocando o * indo até o endereço de memória e lendo que está dentro.
	fmt.Println(variavel3, *ponteiro)
}
