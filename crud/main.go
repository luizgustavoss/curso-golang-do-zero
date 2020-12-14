package main

import (
	"crud/server"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", server.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", server.RecuperarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", server.AlterarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", server.ExcluirUsuario).Methods(http.MethodDelete)
	router.HandleFunc("/usuarios/{id}", server.RecuperarUsuario).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":5000", router))
}
