package enderecos

import "strings"

//TipoDeEndereco retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	//enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(endereco, " ")[0]

	enderecoTemUmTIpoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTIpoValido = true
		}
	}

	if enderecoTemUmTIpoValido {
		return primeiraPalavraDoEndereco
	}

	return "Tipo Inválido"
}
