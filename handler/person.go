package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/basic_API_with_GO_and_postgres/model"
)

type person struct {
	storage Storage
}

func NewPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {

	if m := r.Method; m != http.MethodPost {
		response := NewResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		response := NewResponse(Error, "La persona no tiene una estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Create(&data)

	if err != nil {
		response := NewResponse(Error, "Hubo un proble al insertar la persona", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := NewResponse(Message, "Persona creada", err)
	responseJSON(w, http.StatusOK, response)

}

func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := NewResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return

	}

	resp, err := p.storage.GetAll()

	if err != nil {

		response := NewResponse(Error, "Hubo un proble al obtener  las personas", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return

	}

	response := NewResponse(Message, "SD", resp)
	responseJSON(w, http.StatusOK, response)

}

func (p *person) update(w http.ResponseWriter, r *http.Request) {

	if m := r.Method; m != http.MethodPut {
		response := NewResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || ID < 0 {
		response := NewResponse(Error, "El ID debe ser un número entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		response := NewResponse(Error, "Persona no encontrada", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Update(ID, &data)

	if err != nil {
		response := NewResponse(Error, "Persona no encontrada", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := NewResponse(Message, "Persona actualizada", nil)
	responseJSON(w, http.StatusOK, response)

}

func (p *person) delete(w http.ResponseWriter, r *http.Request) {
	if m := r.Method; m != http.MethodDelete {
		response := NewResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || ID < 0 {
		response := NewResponse(Error, "El ID debe ser un número entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Delete(ID)

	if errors.Is(err, model.IDPersonNotFount) {
		respose := NewResponse(Error, "El reistro no existe", nil)
		responseJSON(w, http.StatusNotFound, respose)
		return
	}

	if err != nil {
		respose := NewResponse(Error, "Algo salio mal al eliminar registro", nil)
		responseJSON(w, http.StatusInternalServerError, respose)
		return
	}
	
	respose := NewResponse(Message, "ok", nil)
	responseJSON(w, http.StatusOK, respose)



}
