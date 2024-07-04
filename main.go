package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Rotas
	router.HandleFunc("/info", getInfoMe).Methods("GET")
	router.HandleFunc("/convert", convertImage).Methods("POST")
	router.HandleFunc("/convert/webp", convertImageWebp).Methods("POST")

	// Inicia o servidor
	fmt.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
