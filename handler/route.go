package handler

import "net/http"

/**
	Esta función lo que hace es crear un enrutador para crear una persona 
	resive un mux, y un objeto que implemente la interfaz storage.

**/
func RoutePerson(mux *http.ServeMux, storage Storage){
	h := NewPerson(storage)//creamos una instancia de persona y le pasamos storage(un objeto que implementa la interfaz)


	mux.HandleFunc("/v1/persons/create", h.create)//registramos el handler dentro del mux y la ruta
	mux.HandleFunc("/v1/persons/getall",h.getAll)//agregamos el endpoint para obtener todos los registros
	mux.HandleFunc("/v1/persons/update",h.update)
	mux.HandleFunc("/v1/persons/delete",h.delete)

}