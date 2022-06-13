package repositorio

import (
	"database/sql"
	"errors"
	"log"
	"main/modelos"
)

//MACAPA INSERIR NO MYSQL
type repositorioMacapa struct {
	db *sql.DB
}

func NewRepositorioMacapa(db *sql.DB) *repositorioMacapa {
	return &repositorioMacapa{db}
}

func (repo repositorioMacapa) CriarListaMacapa(lista []modelos.Contatos) error {
	statement, err := repo.db.Prepare(`insert into contacts(nome, celular)values(?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	defer repo.db.Close()

	for _, v := range lista {
		_, err := statement.Exec(v.Nome, v.Celular)
		if err != nil {
			return errors.New("erro ao adicionar lista de usuario")

		}
		continue
	}

	return nil
}
