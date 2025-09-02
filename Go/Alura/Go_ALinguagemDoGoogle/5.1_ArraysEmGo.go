package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

//pacote para web - responsáveis por requisições

func main() {

	exibeIntroducao()

	for {
		exibeNomes()

		exibeMenu()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1) //Indica que ocorreu alguma coisa inesperada.
		}
	}

}

func exibeIntroducao() {
	nome := "Fabio"
	versao := 1.1
	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	var sites [4]string //desvantagem do array é que tem tamanho pré estabelecido.
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.metanikk.com.br"

	fmt.Println(sites)

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)
	// fmt.Println(resp)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site: ", site, "está com problemas. Status Code: ", resp.StatusCode)
	}
}

func exibeNomes() {
	nomes := []string{"Felipe", "Fabio", "Luana"}
	fmt.Println("O meu slice tem ", len(nomes))
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "items")

	nomes = append(nomes, "Maria", "José")

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem ", len(nomes))
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "items")
	//Slice é uma abstração de um array, ele cria um array por de baixo dos panos já com o tamanho eu irei precisar;
}
