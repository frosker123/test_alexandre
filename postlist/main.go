package main

import (
	"fmt"
	"log"
	"main/keytoken"
	"main/router"
	"net/http"
)

func main() {
	router := router.GerarAPI()
	t1, _ := keytoken.CreateTokenMacapa()
	t2, _ := keytoken.CreateTokenVarejao()
	fmt.Println("token macapa : ", t1)
	fmt.Println("token varejao : ", t2)

	port := ":9000"

	fmt.Println("rodando api ")
	log.Fatal(http.ListenAndServe(port, router))
}
