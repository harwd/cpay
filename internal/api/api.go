package api

import "net/http"

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /payments", createPayment)

	return mux
}

func createPayment(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}