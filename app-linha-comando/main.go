package main

import (
	"fmt"
	"log"
	"os"
	"scrapper/app"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()

	if error := aplicacao.Run(os.Args); error != nil{
		log.Fatal(error)
	}

}
