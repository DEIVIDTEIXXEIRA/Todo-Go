package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var RotaEquipes = []Rota{
	{
		Uri:                "/equipe",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEquipes,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/equipe",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarEquipes,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/equipes/{equipeId}/editar",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoDeEquipe,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/equipes/{equipeId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.EditarEquipe,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/equipes/{equipeId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarEquipe,
		RequerAutenticacao: true,
	},
}
