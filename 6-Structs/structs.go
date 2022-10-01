package main

import "fmt"

type usuario struct{
	nome string
	idade uint8
	endereco endereco
}

type endereco struct{
	logradouro string
	numero uint8
}

func main() {
	fmt.Println("Arquivo struct")

	var u usuario
	u.nome = "Davi"
	u.idade = 12
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua do Mario", 23}
	usuario2 := usuario{"Silvia", 18, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Helena"}
	fmt.Println(usuario3)
}