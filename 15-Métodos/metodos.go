package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func main() {
	usuario1 := usuario{"Silvia", 18}
	usuario1.salvar()

	usuario2 := usuario{"Luiza", 18}
	usuario2.salvar()
}