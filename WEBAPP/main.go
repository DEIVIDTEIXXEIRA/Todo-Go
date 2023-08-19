package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplate()

	fmt.Printf("rodando Webapp %d\n", config.Porta)

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":9090", r))
}
