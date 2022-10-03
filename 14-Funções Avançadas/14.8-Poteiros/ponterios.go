package main

import "fmt"

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	inverterSinalPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
 
func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalPonteiro(numero *int){
	*numero = *numero * -1
}