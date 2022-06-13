package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rotas struct {
	URI    string
	Metodo string
	Funcao func(w http.ResponseWriter, r *http.Request)
}

func ConfiguraRotas(router *mux.Router) *mux.Router {
	rotas := rotasSetadas

	for _, rota := range rotas {
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)

	}
	return router
}
