package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	local := "rua dos loucos"

	localTop := endereco.TipoDeEndereco(local)

	fmt.Println(localTop)
}
