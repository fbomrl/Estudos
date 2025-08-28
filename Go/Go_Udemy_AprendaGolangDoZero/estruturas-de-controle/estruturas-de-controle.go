package main

import "fmt"

func main() {
	fmt.Println("ESTRUTURAS DE CONTROLE")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//Quando você cria uma variável no If vc ta limitando a variável ao escopo do if/else que ele foi criado.
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if numero < -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Número está entre 0 e -10")
	}
}
