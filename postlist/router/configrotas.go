package router

import (
	"main/midllewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Rotas struct {
	URI          string
	Metodo       string
	Funcao       func(w http.ResponseWriter, r *http.Request)
	Autenticacao bool
}

func ConfiguraRotas(router *mux.Router) *mux.Router {
	rotas := configRotas

	for _, rota := range rotas {
		if rota.URI == "/varejao" {
			router.HandleFunc(rota.URI, midllewares.AutenticaVarejao(rota.Funcao)).Methods(rota.Metodo)

		} else {
			router.HandleFunc(rota.URI, midllewares.AutenticaMacapa(rota.Funcao)).Methods(rota.Metodo)

		}

	}

	return router
}
