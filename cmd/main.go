package main

import (
	"log"
	"net/http"

	"github.com/basic_API_with_GO_and_postgres/handler"
	"github.com/basic_API_with_GO_and_postgres/storage"
)

func main() {
	store := storage.NewMemory()

	mux := http.NewServeMux()

	handler.RoutePerson(mux, store)

	log.Println("Servidor iniciado en el puerto :8080")
	
	log.Fatal(http.ListenAndServe(":3000", mux))


}