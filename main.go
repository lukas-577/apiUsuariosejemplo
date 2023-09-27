package main

import (
	"apiusuarios/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/usuarios", handlers.GetUsuariosHandler).Methods("GET")

	http.Handle("/", r)

	// Iniciar el servidor
	http.ListenAndServe(":8080", nil)
}
