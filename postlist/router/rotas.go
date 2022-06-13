package router

import (
	"main/rest"
	"net/http"
)

var configRotas = []Rotas{
	{
		URI:          "/varejao",
		Metodo:       http.MethodPost,
		Funcao:       rest.Varejao,
		Autenticacao: true,
	},
	{
		URI:          "/macapa",
		Metodo:       http.MethodPost,
		Funcao:       rest.Macapa,
		Autenticacao: true,
	},
}
