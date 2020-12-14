package main

import "fmt"

func funcao1(){
	fmt.Println("Função 1")
}

func funcao2(){
	fmt.Println("Função 2")
}

func funcao3(){
	fmt.Println("Função 3")
}
func main() {
	defer funcao1()
	defer funcao2()
	funcao3()
}
