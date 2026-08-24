package api

import (
	"fmt"
	"io"
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /payments", createPayment)

	return mux
}

func createPayment(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading body", http.StatusBadRequest)
		return
	}

	body := string(bytes)
	fmt.Println(body)

	w.WriteHeader(http.StatusCreated)
}