package enderecos

import "strings"

// TipoDeEndereco é responsável por verificar se o tipo de endereço é válido
func TipoDeEndereco(endereco string) string {
	tiposValidaos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoMinusculo := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoMinusculo, " ")[0]
	enderecoTemTipoValido := false
	for _, tipo := range tiposValidaos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemTipoValido = true
		}
	}
	if enderecoTemTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}
	return "Tipo Inválido"
}
