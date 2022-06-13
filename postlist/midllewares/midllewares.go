package midllewares

import (
	"errors"
	"main/keytoken"
	"main/loggs"
	"net/http"
)

func AutenticaMacapa(autoriza http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if valido := keytoken.TokenvalidoMacapa(r); !valido {
			loggs.ResponseError(w, http.StatusBadRequest, errors.New("token inserido macapa não é valido"))
			return
		}
		autoriza(w, r)
	}
}

func AutenticaVarejao(autoriza http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if valido := keytoken.TokenvalidoVarejao(r); !valido {
			loggs.ResponseError(w, http.StatusBadRequest, errors.New("token inserido varejao não é valido"))
			return
		}
		autoriza(w, r)
	}
}
