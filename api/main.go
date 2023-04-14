package main

import (
	"VitoriaXaavier/Alba-Social-Network/src/router"
	"log"
	"net/http"
)

func main() {

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
	
}