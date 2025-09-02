package main

import "fmt"

func main() {
	var n int
	fmt.Println("Digite um número:")
	fmt.Scanf("%d", &n)

	switch {
	case n%3 == 0:
		if n%5 == 0 {
			if n%7 == 0 {
				fmt.Printf("PlingPlangPlong")
				break
			}
			fmt.Printf("PlingPlang")
			break
		}
		if n%7 == 0 {
			fmt.Printf("PlingPlong")
			break
		}
		fmt.Printf("Pling")
	case n%5 == 0:
		if n%7 == 0 {
			fmt.Printf("PlangPlong")
			break
		}
		fmt.Printf("Plang")
	case n%7 == 0:
		fmt.Printf("Plong")
	default:
		fmt.Println(n)
	}
}
