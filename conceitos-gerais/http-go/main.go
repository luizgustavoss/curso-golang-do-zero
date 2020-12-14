package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/ola", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Olá Mundo!"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
	
}
