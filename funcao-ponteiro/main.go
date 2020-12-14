package main

import "fmt"

func inverterSinalComPonteiro(valor *int){
	*valor = *valor * -1
}

func main() {
	valor := 40
	fmt.Println(valor)
	inverterSinalComPonteiro(&valor)
	fmt.Println(valor)
}
