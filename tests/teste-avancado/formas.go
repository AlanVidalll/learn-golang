package formas

import (
	"math"
)

type Forma interface { // interface pode receber varios tipos de dados que exista um metodo area() atrelado a nele
	area() float64
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

type Circulo struct {
	raio float64
}

func (r Retangulo) area() float64 {
	return r.Altura * r.Largura
}

func (c Circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}
