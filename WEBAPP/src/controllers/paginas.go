package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
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

func CarregarPaginaDeEdicaoDeTarefa(w http.ResponseWriter, r* http.Request) {
	parametros:= mux.Vars(r)
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