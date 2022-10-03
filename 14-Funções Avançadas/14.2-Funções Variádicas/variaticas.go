package main

import "fmt"

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 88, 889, 3123)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo", 10, 4, 5, 6, 7, 8)
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros{
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int){
	for _, numero := range numeros{
		fmt.Println(texto, numero)
	}
}