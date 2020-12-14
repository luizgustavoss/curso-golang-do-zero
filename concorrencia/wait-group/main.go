package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func(){
		escrever("Olá Mundo!")
		waitGroup.Done()
	}()

	go func(){
		escrever("Hello World!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string){
	i := 0
	for i < 5{
		i++
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}