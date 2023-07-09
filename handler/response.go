package handler

import (
	"encoding/json"
	"net/http"
)
const(
	Error = "error"
	Message = "message"
)
type response struct {
	MessageType string      `json:"message-type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func NewResponse(messageTipe string, message string, data interface{}) response {
	return response{messageTipe, message, data}
}

func responseJSON(w http.ResponseWriter, ststuscode int, r response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(ststuscode)

	err := json.NewEncoder(w).Encode(&r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
