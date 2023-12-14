package enderecos

import (
	"testing"
)

type cenarioTipoEndereco struct {
	enderecoInformado string
	enderecoEsperado  string
}

func TestTipoEndereco(t *testing.T) {
	cenariosTipoEndereco := []cenarioTipoEndereco{
		{"Rua Teste", "Rua"},
		{"Avenida Teste", "Avenida"},
		{"Estrada Teste", "Estrada"},
		{"Rodovia Teste", "Rodovia"},
		{"Praça Teste", "Tipo Inválido"},
	}

	for _, cenario := range cenariosTipoEndereco {
		tipoEnderecoRetorno := TipoEndereco(cenario.enderecoInformado)

		if tipoEnderecoRetorno != cenario.enderecoEsperado {
			t.Errorf("Tipo Inválido: Esperado '%s' e o retorno foi '%s'", cenario.enderecoEsperado, tipoEnderecoRetorno)
		}
	}
}
