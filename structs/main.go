package main

import "fmt"

type usuario struct {
	nome string
	idade int8
}

func main(){

	var u1 usuario
	u1.nome = "Luiz"
	u1.idade = 37
	fmt.Println(u1)

	u2 := usuario{"Gustavo", 37}
	fmt.Println(u2)

	u3 := usuario{nome: "Luiz"}
	fmt.Println(u3)
}
