package main

import "fmt"

type pessoa struct{
	nome string
	sobrenome string
	idade int16
	altura int16
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {

	var estudante1 estudante
	estudante1.nome = "Luiz"
	estudante1.sobrenome = "Stábile"
	estudante1.idade = 37
	estudante1.altura = 180
	estudante1.curso = "Arquitetura de Software"
	fmt.Println(estudante1)

	var pessoa = pessoa{sobrenome: "Souza", nome:"Gustavo", idade: 37, altura: 180 }
	estudante2 := estudante{pessoa, "Engenharia", "USP"}
	fmt.Println(estudante2)
}
