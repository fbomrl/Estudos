package main
import "fmt"
import "reflect"

func main() {

	nome := "Fabio" //Consegue infeferir a variável, retirando o string, ele já entende que é uma string, etc.
	idade := 36 //inferencia dos tipos de variáveis (int, string, boolean).
	versao := 1.1
	fmt.Println("Olá sr.", nome, "sua idade é ", idade)
	fmt.Println("Este programa está na versão", versao)


	//Verificar se a inferencia automática do float é de 32 ou 64.
	fmt.Println("O tipo da variável versão é ", reflect.TypeOf(versao))

	// 	Aula 2.1 - No exemplo da aula é mostrado que quando você não atribui um valor pra variável (no caso idade) ele fica com 
	// valor default, que nesse caso é zero.

	// 	Aula 2.1 - Outra situação que ocorre é quando você declara uma variável e não a utiliza, assim que você executa o programa
	// ele acusa "idade declared and not used" e assim o código não executa, (ou remove ou utiliza ela)

	//	Aula 2.2 - Ao invés de utilizar o "var" na declaraçãod e variáveis, você pode utilziar o := que vai
	//declarar e inferir a variável;

}