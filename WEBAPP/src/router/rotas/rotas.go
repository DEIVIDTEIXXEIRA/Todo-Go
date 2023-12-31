package rotas

import (
	"net/http"
	"webapp/src/midlewares"

	"github.com/gorilla/mux"
)

type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(router *mux.Router) *mux.Router {
	rotas := RotasLogin
	rotas = append(rotas, RotasUsuarios...)
	rotas = append(rotas, RotaPaginaPrincipal)
	rotas = append(rotas, RotaDeTarefas...)
	rotas = append(rotas, RotaEquipes...)

	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			router.HandleFunc(rota.Uri,
				midlewares.Logger(midlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			router.HandleFunc(rota.Uri, midlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}

		router.HandleFunc(rota.Uri, rota.Funcao).Methods(rota.Metodo)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
