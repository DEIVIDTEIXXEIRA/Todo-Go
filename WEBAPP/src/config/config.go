package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APIURL   = ""
	Porta    = 0
	HashKey  []byte
	BlockKey []byte
)

func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal()
	}

	Porta, erro = strconv.Atoi(os.Getenv("Porta"))
	if erro != nil {
		log.Fatal()
	}

	APIURL = os.Getenv("APIURL")
	HashKey = []byte(os.Getenv("HashKey"))
	BlockKey = []byte(os.Getenv("BlockKey"))

}
