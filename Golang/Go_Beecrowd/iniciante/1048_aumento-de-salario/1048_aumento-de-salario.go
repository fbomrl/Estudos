package main

import "fmt"

func main() {
	var salario float64
	novoSalario := 0.00
	percentual := 0.00
	reajuste := 0.00

	fmt.Scan(&salario)

	if salario > 0 && salario <= 400.00 {
		percentual = 0.15 * 100
		reajuste = salario * 0.15
		novoSalario = reajuste + salario
	}
	if salario > 400.00 && salario <= 800.00 {
		percentual = 0.12 * 100
		reajuste = salario * 0.12
		novoSalario = reajuste + salario
	}
	if salario > 800.00 && salario <= 1200.00 {
		percentual = 0.10 * 100
		reajuste = salario * 0.10
		novoSalario = reajuste + salario
	}
	if salario > 1200.00 && salario <= 2000.00 {
		percentual = 0.07 * 100
		reajuste = salario * 0.07
		novoSalario = reajuste + salario
	}
	if salario > 2000.00 {
		percentual = 0.04 * 100
		reajuste = salario * 0.04
		novoSalario = reajuste + salario
	}
	fmt.Printf("Novo salario: %.2f\n", novoSalario)
	fmt.Printf("Reajuste ganho: %.2f\n", reajuste)
	fmt.Printf("Em percentual: %.f %%\n", percentual)

}
