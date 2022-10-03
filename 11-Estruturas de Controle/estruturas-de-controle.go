package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10

	if numero > 15{
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual que 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número maior que zero")
	} else {
		fmt.Println("Número menor que zero")
	}
}