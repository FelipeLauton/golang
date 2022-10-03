package main

import "fmt"

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(200)
	fmt.Println(dia)

	dia2 := diaDaSemana2(2)
	fmt.Println(dia2)
}

func diaDaSemana(numero int) string{
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça Feira"
	case 4:
		return "Quarta Feira"	
	case 5:
		return "Quinta Feira"	
	case 6:
		return "Sexta Feira"	
	case 7:
		return "Sabádo"		
	default:
		return "Número invalido"
	}
}

func diaDaSemana2(numero int) string {
	diaDaSemana := ""
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "Terça Feira"
	case numero == 4:
		diaDaSemana = "Quarta Feira"	
	case numero == 5:
		diaDaSemana = "Quinta Feira"	
	case numero == 6:
		diaDaSemana = "Sexta Feira"	
	case numero == 7:
		diaDaSemana = "Sabádo"		
	default:
		diaDaSemana = "Número invalido"
	}
	return diaDaSemana 
}