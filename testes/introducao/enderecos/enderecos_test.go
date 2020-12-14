package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {

	cenarios := []cenarioDeTeste {
		{ "Avenida Paulista", "Avenida"},
		{ "RUA Cardoso", "Rua"},
		{ "Rua Zambas", "Rua"},
		{ "Alameda Zambas", "Tipo Inválido"},
		{ "Estrada Água da Onça", "Estrada"},
		{ "Rodovia dos Imigrantes", "Rodovia"},
	}

	for _, cenario := range cenarios {
		tipoDeEnderecoRetornado := TipoDeEndereco(cenario.enderecoInserido)
		if cenario.retornoEsperado != tipoDeEnderecoRetornado {
			t.Errorf("O tipo recebido é diferente do esperado. Esperava %s e recebeu %s.",
				cenario.retornoEsperado, tipoDeEnderecoRetornado)
		}
	}
}


type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado string
}