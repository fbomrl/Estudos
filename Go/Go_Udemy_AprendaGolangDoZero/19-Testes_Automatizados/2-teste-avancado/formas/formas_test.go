package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	//TDD - Test Driven Development (Programação Orientada a Testes)
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			//Fatal faz parar aqui e não da continuidade no teste, para de executar.
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
			// t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()
		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
}
