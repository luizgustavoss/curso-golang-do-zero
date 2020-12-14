package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Canino struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
	Idade uint8 `json:"idade"`
}

func main() {
	marshal()
	unmarshal()
}

func unmarshal(){
	dogString := `{"nome":"Lulu","raca":"SRD","idade":5}`

	var dog Canino

	if error := json.Unmarshal([]byte(dogString), &dog); error != nil {
		log.Fatal(error)
	}
	fmt.Println(dog)
}

func marshal(){
	cachorro := Canino{"Lulu", "SRD", 5}
	jsonDog, error := json.Marshal(cachorro)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(bytes.NewBuffer(jsonDog))
}
