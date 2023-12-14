package main

import (
	"fmt"
	"start-test-go/enderecos"
)

func main() {

	tipoEndereco := enderecos.TipoEndereco("Rua Santo Antonio")
	fmt.Println(tipoEndereco)
}
