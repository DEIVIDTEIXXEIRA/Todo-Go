package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var RotaDeTarefas = []Rota{
	{
		Uri:                "/tarefas",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarTarefa,
		RequerAutenticacao: true,
	},
}
