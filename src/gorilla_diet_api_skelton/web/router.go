package web

import (
	"gorilla_diet_api_skelton/handler"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/", handler.Root).Methods("GET")
	return m
}
