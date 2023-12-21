package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {

	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := CalcArea(ret)

		if areaEsperada != areaRecebida {
			t.Fatalf("Área esperada: %f, difere da área recebeida: %f ", areaEsperada, areaRecebida)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := CalcArea(circ)

		if areaEsperada != areaRecebida {
			t.Fatalf("Área esperada: %f, difere da área recebeida: %f ", areaEsperada, areaRecebida)
		}
	})
}
