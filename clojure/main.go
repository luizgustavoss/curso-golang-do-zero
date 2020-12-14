package main

import "fmt"

func clojure() func(){
	texto := "Dentro da função clojure"

	funcao := func(){
		fmt.Println(texto)
	}
	return funcao
}


func main() {

	texto := "Dentro da função main"
	fmt.Println(texto)

	funcao := clojure()
	funcao()
}
