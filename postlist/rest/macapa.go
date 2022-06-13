package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"main/loggs"
	"main/tipo"
	"net/http"
)

func Macapa(w http.ResponseWriter, r *http.Request) {

	var lista []tipo.Contatos
	const URL = "http://localhost:8080/contatos_macapa"

	body, err := io.ReadAll(r.Body)
	if err != nil {
		err = errors.New("erro ao ler body")
		loggs.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &lista)
	if err != nil {
		err = errors.New("erro ao fazer unmarshal")
		loggs.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	marsha, err := json.Marshal(lista)
	if err != nil {
		err = errors.New("erro ao fazer marshal")
		loggs.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	response, err := http.Post(URL, "application/json", bytes.NewBuffer(marsha))
	if err != nil {
		err = errors.New("erro na request api")
		loggs.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	defer response.Body.Close()

	res, err := io.ReadAll(response.Body)
	if err != nil {
		err = errors.New("erro pra ler o body")
		loggs.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	fmt.Println(string(res))

	if response.StatusCode >= 400 {
		err = errors.New("erro na request api")
		loggs.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	loggs.ResponseJson(w, http.StatusOK, "lista criada")
}
