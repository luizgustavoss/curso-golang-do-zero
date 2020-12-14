package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {

	t.Run("Retângulo", func(t *testing.T){
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRetornada := ret.area()

		if areaEsperada != areaRetornada {
			t.Errorf("Área inválida! Esperado: %0.2f | Retornada: %0.2f ", areaEsperada, areaRetornada)
		}
	})

	t.Run("Circulo", func(t *testing.T){
		cir := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRetornada := cir.area()

		if areaEsperada != areaRetornada {
			t.Errorf("Área inválida! Esperado: %0.2f | Retornada: %0.2f ", areaEsperada, areaRetornada)
		}
	})
}