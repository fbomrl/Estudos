package main

import "fmt"

func main() {
	fmt.Println("MAPS")
	//os tipos de dados sempre tem que ser iguais, não é possível colocar um valor de string e outro de int por exemplo.
	usuario := map[string]string{
		"nome":      "Fabio",
		"sobrenome": "Meireles",
	}
	fmt.Println(usuario)
	fmt.Println("--------------------")

	//possibilidade de incluir um map com tipo map
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Fabio",
			"ultimo":   "Silva",
		},
		"curso": {
			"nome":   "Ciências Biológicas",
			"campus": "Campus II",
		},
	}
	fmt.Println(usuario2)
	fmt.Println("--------------------")
	//possibilidade de deletar uma chave do map
	delete(usuario2, "nome")
	fmt.Println(usuario2)
	fmt.Println("--------------------")
	//possibilidade de incluir uma chave no map
	usuario2["signo"] = map[string]string{
		"nome": "Áries",
	}
	fmt.Println(usuario2)
}
