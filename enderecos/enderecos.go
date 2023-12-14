package enderecos

import (
	"strings"
)

// TipoEndereco recebe um endereço e retorna seu tipo (Rua, Avenida ...)
func TipoEndereco(endereco string) string {
	tiposEnderecoValidos := []string{"rua", "avenida", "rodovia"}

	prefixoEndereco := strings.ToLower(endereco)
	tipoEndereco := strings.Split(prefixoEndereco, " ")[0]

	enderecoValido := false

	for _, tipo := range tiposEnderecoValidos {
		if tipo == tipoEndereco {
			enderecoValido = true
		}
	}

	if enderecoValido {
		return strings.Title(tipoEndereco)
	}

	return "Tipo Inválido"
}
