package main

import "fmt"

func recuperarExecucao(){
	if r:= recover(); r != nil{
		fmt.Println("Execução recuperada!")
	}
}

func verificaAprovacao(nota1, nota2 float64) bool {
	defer recuperarExecucao()
	media := (nota1 + nota2) / 2

	if media > 6{
		return true
	} else if media < 6 {
		return false
	}

	panic("A média precisa ser diferente de 6!")
}


func main() {

	aprovado := verificaAprovacao(6, 6 )

	fmt.Println("Aluno aprovado: ", aprovado)
}
