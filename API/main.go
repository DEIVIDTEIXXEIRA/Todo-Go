package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Gerar()
	fmt.Println("rodando api")

	log.Fatal(http.ListenAndServe(":8080", r))
}
