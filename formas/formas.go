package formas

import (
	"math"
)

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (r Retangulo) area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) area() float64 {
	return math.Pi * math.Pow(c.Raio, 2) //Raio ao quadrado (raio * raio)
}

type Forma interface {
	area() float64
}

//CalcArea calcula a área do objeto
func CalcArea(f Forma) float64 {
	return f.area()
}
