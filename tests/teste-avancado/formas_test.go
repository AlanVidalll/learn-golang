package formas
import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areRecebida := ret.area()

		if areaEsperada != areRecebida {
			t.Errorf("A área recebida %f é diferente da área esperada %f", areRecebida, areaEsperada)
			//t.Fatalf("A área recebida %f é diferente da área esperada %f", areRecebida, areaEsperada) Fatal para de executar assim q da erro, o Error continua executando o resto dos testes.
		}
	})
	t.Run("Circulo", func(t *testing.T) {
		cir := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areRecebida := cir.area()

		if areaEsperada != areRecebida {
			t.Errorf("A área recebida %f é diferente da área esperada %f", areRecebida, areaEsperada)
			//t.Fatalf("A área recebida %f é diferente da área esperada %f", areRecebida, areaEsperada) Fatal para de executar assim q da erro, o Error continua executando o resto dos testes.
		}
	})
}
