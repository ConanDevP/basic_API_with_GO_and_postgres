package handler

import (
	"encoding/json"
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"mesage-type":"error", "message": "Método no permitido"}`))
		return
	}

	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"mesage-type":"error", "message": "La persona no tiene una estructura correcta"}`))
		return
	}

	err = p.storage.Create(&data)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"mesage-type":"error", "message": "Hubo un proble al insertar la persona"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"mesage-type":"message", "message": "Persona creada"}`))

}

func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"mesage-type":"error", "message": "Método no permitido"}`))
		return
	}

	resp, err := p.storage.GetAll()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"mesage-type":"error", "message": "Hubo un proble al obtener  las personas"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"mesage-type":"error", "message": "Hubo un proble al convertir el slice en JSON"}`))
		return
	}

}

func (p *person) update(w http.ResponseWriter, r *http.Request) {

	if m := r.Method; m != http.MethodPut {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"mesage-type":"error", "message": "Método no permitido"}`))
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || ID < 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"mesage-type":"error", "message": "El ID debe ser un número entero positivo"}`))
		return
	}

	data := model.Person{}

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"mesage-type":"error", "message": "Persona no encontrada"}`))
		return
	}

	err = p.storage.Update(ID, &data)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"mesage-type":"error", "message": "Persona no encontrada"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"mesage-type":"ok", "message": "Persona actualizada"}`))
	

}
