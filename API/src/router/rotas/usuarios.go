package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasDeUsuarios = []Rota{
	{
		Uri:                   "/usuarios",
		Metodo:                http.MethodPost,
		Funcao:                controllers.CriarUsuario,
		RequerAutententicacao: false,
	},
	{
		Uri:                   "/usuarios",
		Metodo:                http.MethodGet,
		Funcao:                controllers.BuscarUsuarios,
		RequerAutententicacao: false,
	},
	{
		Uri:                   "/usuarios/{usuarioId}",
		Metodo:                http.MethodGet,
		Funcao:                controllers.BuscarUsuario,
		RequerAutententicacao: false,
	},
	{
		Uri:                   "/usuarios/{usuarioId}",
		Metodo:                http.MethodPut,
		Funcao:                controllers.EditarUsuario,
		RequerAutententicacao: false,
	},
	{
		Uri:                   "/usuarios/{usuarioId}",
		Metodo:                http.MethodDelete,
		Funcao:                controllers.DeletarUsuario,
		RequerAutententicacao: false,
	},
}
