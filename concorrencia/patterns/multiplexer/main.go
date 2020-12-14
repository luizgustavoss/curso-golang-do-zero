package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	canal := multiplexar(escrever("Olá Mundo"), escrever("Hello World"))

	for i := 0; i< 10 ; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(entrada1, entrada2 <-chan string) <-chan string {

	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <- entrada1:
				canalSaida <- mensagem
			case mensagem := <- entrada2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida
}

func escrever(texto string) <-chan string {
	canal:= make(chan string)

	go func(){
		for{
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}