package main

import (
	"errors"
	"fmt"
)

func main() {
	//--NÚMEROS INTEIROS--
	var numero int64 = -1000000000
	fmt.Println(numero)

	var numero2 uint = 10000
	fmt.Println(numero2)

	//INT32 = RUNE
	var numero3 rune = 12456
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	//--NÚMEROS REAIS--

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.4
	fmt.Println(numeroReal3)

	//--STRINGS--
	
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	//--BOOLEAN--

	var booleano1 bool = true
	fmt.Println(booleano1)

	//--ERROR--
	
	var erro error
	fmt.Println(erro)

	var erroPersonalizado error = errors.New("Erro interno")
	fmt.Println(erroPersonalizado)

}
