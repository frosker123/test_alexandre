package router

import (
	"main/rest"
	"net/http"
)

var rotasSetadas = []Rotas{
	{
		URI:    "/contatos_varejao",
		Metodo: http.MethodPost,
		Funcao: rest.PostContatosVarejao,
	},
	{
		URI:    "/contatos_macapa",
		Metodo: http.MethodPost,
		Funcao: rest.PostContatosMacapa,
	},
}
