package main

import (
	"fmt"
	"math"
)

type forma interface { // interface pode receber varios tipos de dados que exista um metodo area() atrelado a nele
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

func escreverArea(f forma) {
	fmt.Printf("A area da forma e %0.2f\n", f.area())
}

func main() {

	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{100}
	escreverArea(c)
}
