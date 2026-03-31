package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HellowWorldResponse struct{
	Message string `json:"message"`
}

func (h *Handler)HomeHandler(w http.ResponseWriter, r *http.Request){
	response := HellowWorldResponse{
		Message: "Hello world",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler)DBCheckHandler(w http.ResponseWriter, r *http.Request){
	err := h.DB.Ping()
	if err != nil {
		http.Error(w, "Database tidak terhubung", http.StatusInternalServerError)
		return
	}
	response := HellowWorldResponse{
		Message: "DB connected",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func (h *Handler)EchoHandler(w http.ResponseWriter, r *http.Request){
	var body map[string]any

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil{
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}

	fmt.Println("Body input", body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(body)
}
