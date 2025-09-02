package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaUm := ContaCorrente{"Fabio", 1414, 141414, 9500.00}
	contaDois := ContaCorrente{"Luana", 1414, 131313, 1500.00}

	fmt.Println(contaUm, contaDois)

	var contaTres *ContaCorrente
	contaTres = new(ContaCorrente)
	contaTres.titular = "Felipe"

	fmt.Println(*contaTres)
}
