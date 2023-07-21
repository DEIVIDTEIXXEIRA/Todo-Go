package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa o formato das rotas
type Rota struct {
	Uri                   string
	Metodo                string
	Funcao                func(http.ResponseWriter, *http.Request)
	RequerAutententicacao bool
}

//Configurar coloca todas as rotas dentro do router ja configurado 
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasDeUsuarios
	rotas = append(rotas, rotaLogin)

	for _, rota := range rotas {
		if rota.RequerAutententicacao {
			r.HandleFunc(rota.Uri, middlewares.Autenticar(rota.Funcao)).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.Uri, rota.Funcao).Methods(rota.Metodo)
		}
	}

	return r
	
}
