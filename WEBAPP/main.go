package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Carregar()
	utils.CarregarTemplate()

	fmt.Printf("rodando Webapp %d\n", config.Porta)

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":9090", r))
}
