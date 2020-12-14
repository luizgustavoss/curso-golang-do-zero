package main

import (
	"fmt"
)

func main() {

	// sem um buffer o canal bloquearia até que alguém pudesse ler
	// com buffer, enquanto a capacidade máxima não for atingida, o código segue executando
	canal := make(chan string, 2);
	canal <- "Olá Mundo!"

	mensagem1 := <-canal

	fmt.Println(mensagem1)
}
