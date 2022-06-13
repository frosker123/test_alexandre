package conectdb

import (
	"database/sql"
	"fmt"
	"log"
	"main/configvariavel"

	_ "github.com/lib/pq"
)

func ConectaDBPsql() (*sql.DB, error) {

	conect := fmt.Sprintf("host=%s port=%d user=%s password=%v  sslmode=disable ", configvariavel.HOST, configvariavel.DB_PORTPSQL, configvariavel.DB_USER, configvariavel.PASSWORD)

	db, err := sql.Open("postgres", conect)
	if err != nil {
		log.Fatal("erro ao conetar")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("erro no ping")
	}

	fmt.Printf("conectou-se ao banco de dados postgres no docker :) \n")

	return db, nil
}
