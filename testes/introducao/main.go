package main

import (
	"fmt"
	"testes/enderecos"
)

func main() {

	tipoEndereco := enderecos.TipoDeEndereco("Rua Alvarenga Peixoto")
	fmt.Println(tipoEndereco)
}
