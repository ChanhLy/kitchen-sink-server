package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	data     any
	httpCode int
}

func SendResponse(w http.ResponseWriter, data any, httpCode int) {
	w.WriteHeader(httpCode)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("Error writing response:", err)
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
	}
}
