package loggs

import (
	"encoding/json"
	"log"
	"net/http"
)

type Erros struct {
	Erro string `json:"Erros"`
}

func ResponseJson(w http.ResponseWriter, statuscode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)

	if dados != nil {
		if err := json.NewEncoder(w).Encode(dados); err != nil {
			log.Fatal("erro no encode")
		}
	}
}

func ResponseError(w http.ResponseWriter, statuscode int, err error) {
	ResponseJson(w, statuscode, Erros{
		Erro: err.Error(),
	})
}
