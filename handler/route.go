package handler

import (
	"net/http"

	"github.com/basic_API_with_GO_and_postgres/middleware"
)

/**
	Esta función lo que hace es crear un enrutador para crear una persona
	resive un mux, y un objeto que implemente la interfaz storage.

**/
func RoutePerson(mux *http.ServeMux, storage Storage){
	h := NewPerson(storage)//creamos una instancia de persona y le pasamos storage(un objeto que implementa la interfaz)


	mux.HandleFunc("/v1/persons/create",middleware.Log(middleware.Authenticatos(h.create)))//registramos el handler dentro del mux y la ruta
	mux.HandleFunc("/v1/persons/getall",middleware.Log(h.getAll))//agregamos el endpoint para obtener todos los registros
	mux.HandleFunc("/v1/persons/update",middleware.Log(h.update))
	mux.HandleFunc("/v1/persons/delete",middleware.Log(h.delete))

}