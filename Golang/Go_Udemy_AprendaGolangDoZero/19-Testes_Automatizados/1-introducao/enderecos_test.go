package endereco

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "rua"},
		{"Avenida Paulista", "avenida"},
		{"Rodovia dos Imigrantes", "rodovia"},
		{"Praça das Rosas", "TIPO INVÁLIDO"},
		{"Estrada Qualquer", "estrada"},
		{"RUA DOS BOBOS", "rua"},
		{"AVENIDA REBOUÇAS", "avenida"},
		{"", "TIPO INVÁLIDO"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo de endereço %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("Teste quebrou")
	}
}
