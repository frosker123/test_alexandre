package configvariavel

import (
	"log"
	"os"
	"strconv"

	"github.com/subosito/gotenv"
)

var (
	DB_PORTPSQL = 0
	DB_PORTMSQL = 0
	API_PORT    = 0
	PASSWORD    = ""
	HOST        = ""
	HOSTMYSQL   = ""
	DB_USER     = ""
	DB_NAME     = ""
)

func VariveisAm() {
	var err error

	err = gotenv.Load("./.env")
	if err != nil {
		log.Fatal("erro ao carregar as variaveis")
	}

	DB_PORTPSQL, err = strconv.Atoi(os.Getenv("DB_PORTPSQL"))
	if err != nil {
		log.Fatal("erro ao converter variavel dbporta")
	}

	DB_PORTMSQL, err = strconv.Atoi(os.Getenv("DB_PORTMSQL"))
	if err != nil {
		log.Fatal("erro ao converter variavel dbporta")
	}

	API_PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatal("erro ao converter variavel apiporta")
	}

	DB_USER = os.Getenv("DB_USER")
	HOST = os.Getenv("HOST")
	PASSWORD = os.Getenv("PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	HOSTMYSQL = os.Getenv("HOSTMYSLQ")
}
