package main

import (
	"fmt"
	"os"
)

type Record struct {
	Day      int
	Amount   float64
	Category string
}

type DaysPeriod struct {
	From int
	To   int
}

func main() {

	records := []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 12, Amount: 10, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}
	Inicio()
	for {
		fmt.Println()
		ExibeMenu()
		var opcao int
		fmt.Scan(&opcao)
		switch opcao {
		case 0:
			fmt.Println("SAINDO...")
			os.Exit(0)
		case 1:
			var diaEscolhido int
			fmt.Println("Digite um dia: ")
			fmt.Scan(&diaEscolhido)
			filtrados := Filter(records, func(r Record) bool {
				return r.Day == diaEscolhido
			})
			if diaEscolhido <= 0 || diaEscolhido > 31 {
				fmt.Println("Digite um valor válido")
			}
			fmt.Println(filtrados)
		case 2:
			var periodoFrom int
			var periodoTo int
			fmt.Println("Digite o dia inicial:")
			fmt.Scan(&periodoFrom)
			fmt.Println()
			fmt.Println("Digite a data final")
			fmt.Scan(&periodoTo)

			resultadoPeriodo := Filter(records, ByDaysPeriod(DaysPeriod{periodoFrom, periodoTo}))
			fmt.Println(resultadoPeriodo)
		case 3:
			var porCategoria string
			fmt.Println("Digite a categoria: ")
			fmt.Scan(&porCategoria)
			fmt.Println()

			resultadoCategoria := Filter(records, ByCategory(porCategoria))
			if len(resultadoCategoria) == 0 {
				fmt.Println("Nenhuma categoria encontrada")
			} else {
				fmt.Println(resultadoCategoria)
			}
		case 4:
			var periodoFrom int
			var periodoTo int
			fmt.Println("Digite o dia inicial:")
			fmt.Scan(&periodoFrom)
			fmt.Println()
			fmt.Println("Digite a data final")
			fmt.Scan(&periodoTo)

			resultadoTotalPeriodo := TotalByPeriod(records, DaysPeriod{periodoFrom, periodoTo})
			fmt.Println("$", resultadoTotalPeriodo)
		case 5:
			var periodoFrom int
			var periodoTo int
			var categoria string
			fmt.Println("Digite o dia inicial:")
			fmt.Scan(&periodoFrom)
			fmt.Println()
			fmt.Println("Digite a data final")
			fmt.Scan(&periodoTo)
			fmt.Println("Digite a categoria: ")
			fmt.Scan(&categoria)
			fmt.Println()

			resultadoCategoriaPeriodo := CategoryExpenses(records, DaysPeriod{periodoFrom, periodoTo}, categoria)

			if len(categoria) == 0 {
				fmt.Println("Nenhuma categoria encontrada")
			} else {
				fmt.Println("Total da categoria", categoria, ", R$", resultadoCategoriaPeriodo)
			}
		default:
			fmt.Print("Digite um valor válido")
		}

	}

}

func Filter(records []Record, f func(Record) bool) []Record {
	var result []Record

	for _, r := range records {
		if f(r) {
			result = append(result, r)
		}
	}
	return result
}
func ByDaysPeriod(day DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= day.From && r.Day <= day.To
	}
}
func ByCategory(categoria string) func(Record) bool {
	return func(record Record) bool {
		return record.Category == categoria
	}
}
func TotalByPeriod(records []Record, day DaysPeriod) float64 {
	var total float64
	for _, r := range records {
		if r.Day >= day.From && r.Day <= day.To {
			total += r.Amount
		}
	}
	return total
}

func CategoryExpenses(records []Record, day DaysPeriod, categoria string) float64 {
	var total float64

	resultadoCategoria := Filter(records, ByCategory(categoria))

	for _, r := range resultadoCategoria {
		if r.Day >= day.From && r.Day <= day.To {
			total += r.Amount
		}
	}
	return total
}

func Inicio() {
	fmt.Println("------------------------------------")
	fmt.Println("SELECIONE UM FILTRO")
	fmt.Println("------------------------------------")
	fmt.Println()
}
func ExibeMenu() {
	fmt.Println("1 - DATA")
	fmt.Println("2 - PERIODO")
	fmt.Println("3 - CATEGORIA")
	fmt.Println("4 - DESPESAS POR PERIODO")
	fmt.Println("5 - DESPESAS TOTAIS POR CATEGORIA")
	fmt.Println("0 - SAIR DO PROGRAMA")
}
