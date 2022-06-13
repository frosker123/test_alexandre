package conectdb

import (
	"database/sql"
	"fmt"
	"log"
	"main/configvariavel"

	_ "github.com/go-sql-driver/mysql"
)

func ConectaDBMsql() (*sql.DB, error) {

	conectMysql := fmt.Sprintf("%v:%v@%v/%v", configvariavel.DB_USER, configvariavel.PASSWORD, configvariavel.HOSTMYSQL, configvariavel.DB_NAME)

	db, err := sql.Open("mysql", conectMysql)
	if err != nil {
		log.Fatal("erro nas variaveis")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("erro no ping")
	}

	fmt.Printf("conectou-se ao banco de dados mysql no docker :) \n")

	return db, nil
}
