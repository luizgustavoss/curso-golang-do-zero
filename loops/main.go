package main

import (
	"fmt"
	"time"
)

func main() {

	i := 0
	for i < 10{
		i++
		time.Sleep(time.Second)
		fmt.Println("Incrementando i", i)
	}

	for j := 0; j < 10; j++{
		time.Sleep(time.Second)
		fmt.Println("Incrementando j", j)
	}

	nomes := [3]string{"João", "Lucas", "Antônio"}

	for _, nome := range nomes{
		fmt.Println(nome)
	}

	for _, letra := range "PALAVRA" {
		fmt.Println(string(letra))
	}

	usuarios := map[string]string {
		"Luiz" : "Gustavo",
		"Mariana" : "Budiski",
	}

	for nome, sobrenome := range usuarios {
		fmt.Println(nome, sobrenome)
	}

}
