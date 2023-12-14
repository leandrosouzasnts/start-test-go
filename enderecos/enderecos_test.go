package enderecos

import (
	"testing"
)

func TestTipoEndereco(t *testing.T) {
	tipoRetorno := TipoEndereco("Avenida Saudade")

	tipoEsperado := "Avenida"

	if tipoEsperado != tipoRetorno {
		t.Error("Tipo do Retorno diverge do Esperado")
	}
}
