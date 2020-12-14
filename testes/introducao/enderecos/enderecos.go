package enderecos

import "strings"

// TipoDeEndereco
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoMinusculo := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enderecoMinusculo, " ")[0]

	enderecoDeTipoValido := false
	for _, tipo := range tiposValidos{
		if tipo == primeiraPalavra {
			enderecoDeTipoValido = true
			break
		}
	}

	if enderecoDeTipoValido {
		return strings.Title(primeiraPalavra)
	}

	return "Tipo Inválido"
}