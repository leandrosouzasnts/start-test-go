package enderecos_test

import (
	. "start-test-go/enderecos" // . é um Alias assim não preciso definir o pacote na hora de usar.
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
		//{"Praça Teste", "Tipo Inválido"},
	}

	for _, cenario := range cenariosTipoEndereco {
		tipoEnderecoRetorno := TipoEndereco(cenario.enderecoInformado)

		if tipoEnderecoRetorno != cenario.enderecoEsperado {
			t.Errorf("Tipo Inválido: Esperado '%s' e o retorno foi '%s'", cenario.enderecoEsperado, tipoEnderecoRetorno)
		}
	}
}

func TestQualquer(t *testing.T) {
	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}
}
