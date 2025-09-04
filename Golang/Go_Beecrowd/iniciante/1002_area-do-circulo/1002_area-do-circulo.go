package main

import "fmt"

func main() {
	n := 3.14159
	var raio float64

	fmt.Scan(&raio)
	area := n * raio * raio
	fmt.Printf("A=%.4f\n", area)
}
