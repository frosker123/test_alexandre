package main

import (
	"fmt"
	"log"
	"main/configvariavel"
	"main/router.go"
	"net/http"
)

func main() {
	subirApi := router.GerarAPI()
	configvariavel.VariveisAm()

	fmt.Println("rodando api...")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", configvariavel.API_PORT), subirApi))

}
