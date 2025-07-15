package main

import "fmt"
import "os"
import "net/http" //pacote para web - responsáveis por requisições

func main(){
	
	exibeIntroducao()
	
	for {
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

func exibeIntroducao(){
	nome := "Fabio"
	versao := 1.1
	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu(){
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func lerComando() int{
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido;
}

func iniciarMonitoramento(){
	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)
	// fmt.Println(resp)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!" )
	}
	else {
		fmt.Println("Site: ", site, "está com problemas. Status Code: ", resp.StatusCode)
	}
}

// // Uma opção alternativa que você pode usar é o seguinte serviço:
// // Httpbin
// // Por meio do Httpbin, é possível simular o status HTTP ao passar o código de status correspondente:

// // https://httpbin.org/status/200
// // https://httpbin.org/status/404
// func iniciarMonitoramento____() {
//     fmt.Println("Monitorando...")
//     site com URL inexistente
//     site := "https://httpbin.org/status/404" // ou 200
//     resp, _ := http.Get(site)

//     if resp.StatusCode == 200 {
//         fmt.Println("Site:", site, "foi carregado com sucesso!")
//     } else {
//         fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
//     }
// }
