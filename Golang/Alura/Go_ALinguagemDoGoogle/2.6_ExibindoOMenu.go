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
	// Scanf precisa indicar que espera um int, mesmo o sistema reconhecendo qual é o tipo da variável
	// fmt.Scanf("%d", &comando) 
	// Já com o Scan não precisa especificar que está esperando o tipo da variável.
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	// & É o endereço/ponteiro da variável;
	// %d É um modificador para indicar que está sendo esperado um número inteiro; (%s - string / %f flutuante)

}