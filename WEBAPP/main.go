package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	fmt.Println("rodando Webapp")
	utils.CarregarTemplate()

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":9090", r))
}
