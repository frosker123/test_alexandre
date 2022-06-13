package repositorio

import (
	"database/sql"
	"errors"
	"main/modelos"
)

//VAREJA0 INSERIR NO POSTGRES
type repositorioVarejao struct {
	db *sql.DB
}

func NewRepositorioVarejao(db *sql.DB) *repositorioVarejao {
	return &repositorioVarejao{db}
}

func (repo repositorioVarejao) CriarListaVarejao(lista []modelos.Contatos) error {
	statement := `insert into varejao.contacts(nome, celular)values($1, $2)`

	defer repo.db.Close()

	for _, v := range lista {
		_, err := repo.db.Exec(statement, v.Nome, v.Celular)
		if err != nil {
			return errors.New("erro ao adicionar lista de numeros")
		}
	}
	return nil
}
