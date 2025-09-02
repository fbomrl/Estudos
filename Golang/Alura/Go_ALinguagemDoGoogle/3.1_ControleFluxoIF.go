package main

import "fmt"

func main(){

	nome := "Fabio"
	versao := 1.1
	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int

	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// }else if comando == 2{
	// 	fmt.Println("Exibindo logs...")
	// }else if comando == 0{
	// 	fmt.Println("Saindo do programa")
	// }else{
	// 	fmt.Println("Não conheço este comando")
	// }

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case:2
		fmt.Println("Exibindo logs...")
	case:0
		fmt.Println("Saindo do programa")
	default:
		fmt.Println("Não conheço este comando")
	}


	//Não tem o Break, não precisa, já que ele entra no valor que é requerido e para.
}