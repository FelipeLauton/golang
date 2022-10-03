package main

import "fmt"

func main() {

	defer funcao1()
	// DEFER = ADIAR
	funcao2()

	fmt.Println(alunoAprovado(7, 8))
}

func funcao1() {
	fmt.Println("Chamando a função 1")
}

func funcao2() {
	fmt.Println("Chamando a função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}