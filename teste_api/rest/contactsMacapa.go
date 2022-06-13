package rest

import (
	"encoding/json"
	"errors"
	"io"
	"main/loggs"
	"main/modelos"
	"main/repositorio"
	conectdb "main/services"
	"main/validar"
	"net/http"
)

func PostContatosMacapa(w http.ResponseWriter, r *http.Request) {
	var criarLista []modelos.Contatos

	body, err := io.ReadAll(r.Body)
	if err != nil {
		err = errors.New("erro ao ler body ")
		loggs.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	err = json.Unmarshal(body, &criarLista)
	if err != nil {
		err = errors.New("erro ao fazer o unmarshal")
		loggs.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	db, err := conectdb.ConectaDBMsql()
	if err != nil {
		err = errors.New("erro ao conectar no banco")
		loggs.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	err = validar.ValidarCampos(criarLista)
	if err != nil {
		err = errors.New("numero invalido ou campo nome não preenchido")
		loggs.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	repo := repositorio.NewRepositorioMacapa(db)
	err = repo.CriarListaMacapa(criarLista)
	if err != nil {
		err = errors.New("erro ao criar lista de contatos")
		loggs.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	loggs.ResponseJson(w, http.StatusOK, criarLista)

}
