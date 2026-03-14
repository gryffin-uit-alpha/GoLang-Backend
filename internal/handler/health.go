package handlers

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func HandlerHealth(w http.ResponseWriter, r *http.Request) {
	data := ResponseData{
		Message: "Server is running well from private package!",
		Status:  200,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
