package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulos/auxiliar"
)

func main() {
	fmt.Println("Hello brave new world!")
	auxiliar.Escrever()

	var error = checkmail.ValidateFormat("luiz@gmail.com")
	fmt.Println(error)
}
