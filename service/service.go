package service

import (
	"encoding/json"
	"net/http"
)

type errorResponce struct {
	Error   string
	Comment string
}

func ErrorResponce(err error, comment string, code int, w http.ResponseWriter) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	EvenError := errorResponce{
		Error:   err.Error(),
		Comment: comment,
	}

	if err := json.NewEncoder(w).Encode(EvenError); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
