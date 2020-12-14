package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá Mundo!")
	escrever("Hello World!")
}

func escrever(texto string){
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}