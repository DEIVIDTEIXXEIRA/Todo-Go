package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}
	utils.ExecutarTemplete(w, "login.html", nil)
}

func CarregarPaginaDeCadastroDeusuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplete(w, "cadastro.html", nil)
}

func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/tarefas", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.Erro{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var Tarefas []modelos.Tarefas
	if erro = json.NewDecoder(response.Body).Decode(&Tarefas); erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.Erro{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplete(w, "home.html", Tarefas)
}

func CarregarPaginaDeEdicaoDeTarefa(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	tarefaId, erro := strconv.ParseUint(parametros["tarefaId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.Erro{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/tarefas/%d", config.APIURL, tarefaId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.Erro{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var tarefa modelos.Tarefas

	if erro = json.NewDecoder(response.Body).Decode(&tarefa); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.Erro{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplete(w, "editar-tarefa.html", tarefa)
}

func CarregarPerfilDoUsuario(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	usuario, erro := modelos.BuscarUsuarioCompleto(usuarioId, r)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.Erro{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplete(w, "perfil.html", usuario)
}

//CarregarPaginaDeEdicaoDoUsuario carrega a página de formulário de edição.
func CarregarPaginaDeEdicaoDoUsuario(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	//canal foi criado, pois, já temos uma função que busca as informações do usuario.
	canal := make(chan modelos.Usuario)
	go modelos.BuscaDadosUsuario(canal, usuarioId, r) 
	usuario:= <-canal 

	if usuario.Id == 0 {
		respostas.JSON(w, http.StatusInternalServerError, respostas.Erro{Erro:"Erro ao buscar usuário"})
		return
	}
	
	utils.ExecutarTemplete(w, "editar-usuario.html", usuario)

}