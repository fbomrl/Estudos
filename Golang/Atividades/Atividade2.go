package main

import "fmt"

func main() {
	var n int
	fmt.Println("Digite um número:")
	fmt.Scanf("%d", &n)

	var resultado string

	if n%3 == 0 {
		resultado += "Pling"
	}
	if n%5 == 0 {
		resultado += "Plang"
	}
	if n%7 == 0 {
		resultado += "Plong"
	}
	if resultado == "" {
		fmt.Println(n)
	}

	fmt.Println(resultado)
}
