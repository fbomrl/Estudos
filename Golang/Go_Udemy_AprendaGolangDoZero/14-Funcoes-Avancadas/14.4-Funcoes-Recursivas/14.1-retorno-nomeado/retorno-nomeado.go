package main

import "fmt"

func calculoMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculoMatematicos(10, 20)

	fmt.Println(soma, subtracao)
}
