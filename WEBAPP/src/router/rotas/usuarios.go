package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var RotasUsuarios = []Rota{
	{
		Uri:                "/criar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeCadastroDeusuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/perfil",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPerfilDoUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/editar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoDoUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/editar-usuario",
		Metodo:             http.MethodPut,
		Funcao:             controllers.EditarUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/editar-senha",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoDoSenha,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/editar-senha",
		Metodo:             http.MethodPost,
		Funcao:             controllers.EditarSenha,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/deletar-usuario",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: true,
	},
}
