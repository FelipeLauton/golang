package main

import "fmt"

func main() {
	//--ARITIMETICOS--
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma ,subtracao, divisao, multiplicacao, restoDivisao)

	//--ATRIBUIÇÃO--
	var variavel string = "string"
	variavel2 := "string 2"
	fmt.Println(variavel, variavel2)

	//--OPERADORES RELACIONAIS--
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//--OPERADORES LÓGICOS--
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!false)

	//--OPERADORES UNÁRIOS--
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20
	fmt.Println(numero)

	numero *= 2
	numero /=10
	numero %= 2
	fmt.Println(numero)
}