package main

import (
	"fmt"
)

type usuario struct{
	nome string
	idade int
}

func (u usuario) salvar(){
	fmt.Printf("Salvando os dados do usuário %s \n", u.nome)
}

func (u *usuario) aumentarIdade(){
	u.idade++
}

func main() {

	usuario := usuario{nome: "Luiz", idade: 37}
	fmt.Println(usuario)
	usuario.aumentarIdade()
	usuario.salvar()
	fmt.Println(usuario)
}
