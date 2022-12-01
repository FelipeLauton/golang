package main

import "fmt"

// A comparação de arrays é possivel quando ambos são do mesmo tipo
// Caso os valores de ambos estejam na mesma ordem o retorno será true na comparação
// O tamanho de um array é uma parte inseparavél do seu tipo. 
// [3]int e [2]int são arrays de tipos diferentes por conta do tamanho. 
func main() {
	blue := [3]int{6, 9, 3}
	red := [3]int{6, 9, 3}

	fmt.Println(blue == red)
}