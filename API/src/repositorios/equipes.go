package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Equipe struct {
	db *sql.DB
}

func NovoRepositorioDeEquipes(db *sql.DB) *Equipe {
	return &Equipe{db}
}

func (repositorio Equipe) CriarEquipe(equipe modelos.Equipes) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into equipes (nome, descricao, autor_id) value(?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(equipe.Nome, equipe.Descricao, equipe.AutorId)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Equipe) Buscar(nomeDaEquipe string) ([]modelos.Equipes, error) {
	nomeDaEquipe = fmt.Sprintf("%%%s%%", nomeDaEquipe)

	linhas, erro := repositorio.db.Query(
		"select id, nome, descricao, autor_id from equipes where nome LIKE ?",
		nomeDaEquipe,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var equipes []modelos.Equipes
	for linhas.Next() {
		var equipe modelos.Equipes

		if erro = linhas.Scan(
			&equipe.Id,
			&equipe.Nome,
			&equipe.Descricao,
			&equipe.AutorId,
		); erro != nil {
			return nil, erro
		}
		equipes = append(equipes, equipe)
	}

	return equipes, nil
}

func (repositorio Equipe) BuscarPorId(equipeId uint64) (modelos.Equipes, error) {
	linha, erro := repositorio.db.Query("select id, nome, descricao, autor_id from equipes where id = ?", equipeId)
	if erro != nil {
		return modelos.Equipes{}, erro
	}
	defer linha.Close()

	var Equipe modelos.Equipes

	if linha.Next() {
		if erro = linha.Scan(
			&Equipe.Id,
			&Equipe.Nome,
			&Equipe.Descricao,
			&Equipe.AutorId,
		); erro != nil {
			return modelos.Equipes{}, erro
		}
	}

	return Equipe, nil
}

func (repositorio Equipe) BuscarEquipesUsuario(usuarioId uint64) ([]modelos.Equipes, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, descricao, autor_id from equipes where autor_id = ?",
		usuarioId,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var equipes []modelos.Equipes
	for linhas.Next() {
		var equipe modelos.Equipes

		if erro = linhas.Scan(
			&equipe.Id,
			&equipe.Nome,
			&equipe.Descricao,
			&equipe.AutorId,
		); erro != nil {
			return nil, erro
		}
		equipes = append(equipes, equipe)
	}

	return equipes, nil
}

func (repositorio Equipe) AtualizarEquipe(equipeId uint64, Equipe modelos.Equipes) error {
	statement, erro := repositorio.db.Prepare("update equipes set nome = ?, descricao = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(Equipe.Nome, Equipe.Descricao, equipeId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Equipe) DeletarEquipe(equipeId uint64) error {
	statement, erro := repositorio.db.Prepare("delete from equipes where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(equipeId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Equipe) CriarTarefaDeEquipe(tarefa modelos.Tarefas, equipeId uint64, usuarioId uint64) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into tarefas_equipe (tarefa, observacao, prazo, autor_id, equipes_id) value(?, ?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(tarefa.Tarefa, tarefa.Obsevacao, tarefa.Prazo, tarefa.AutorId, equipeId)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Equipe) BuscarTarefasDaEquipe(equipeId uint64) ([]modelos.Tarefas, error) {
	linhas, erro := repositorio.db.Query("select id, autor_id, tarefa, observacao, prazo FROM tarefas_equipe WHERE equipes_id = ?", equipeId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var tarefasDaEquipe []modelos.Tarefas

	for linhas.Next() {
		var tarefa modelos.Tarefas

		if erro = linhas.Scan(
			&tarefa.Id,
			&tarefa.AutorId,
			&tarefa.Tarefa,
			&tarefa.Obsevacao,
			&tarefa.Prazo,
		); erro != nil {
			return nil, erro
		}

		tarefasDaEquipe = append(tarefasDaEquipe, tarefa)
	}

	return tarefasDaEquipe, nil
}

func (repositorio Equipe) BuscarTarefaDaEquipe(tarefaId uint64) (modelos.Tarefas, error) {
	linha, erro := repositorio.db.Query("select id, autor_id, tarefa, observacao, prazo from tarefas_equipe WHERE id = ?",
		tarefaId,
	)
	if erro != nil {
		return modelos.Tarefas{}, erro
	}
	defer linha.Close()

	var tarefa modelos.Tarefas

	if linha.Next() {
		if erro = linha.Scan(
			&tarefa.Id,
			&tarefa.AutorId,
			&tarefa.Tarefa,
			&tarefa.Obsevacao,
			&tarefa.Prazo,
		); erro != nil {
			return modelos.Tarefas{}, erro
		}
	}

	return tarefa, nil
}

func (repositorio Equipe) EditarTarefaDaEquipe(tarefaId uint64, Tarefa modelos.Tarefas) error {
	statement, erro := repositorio.db.Prepare(
		"update tarefas_equipe set tarefa = ?, observacao = ?, prazo = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(Tarefa.Tarefa, Tarefa.Obsevacao, Tarefa.Prazo, tarefaId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Equipe) DeletarTarefaDaEquipe(equipeId, tarefaId uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from tarefas_equipe where equipes_id = ? and id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(equipeId, tarefaId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Equipe) Adicionar(equipeId, usuarioId uint64) error {

	var usuarioNick string
	erro := repositorio.db.QueryRow(
		"SELECT nick FROM usuarios WHERE id = ?",
		usuarioId,
	).Scan(&usuarioNick)
	if erro != nil {
		return erro
	}

	statement, erro := repositorio.db.Prepare(
		"insert into usuarios_equipe (equipes_id, usuario_id, usuario_nick) value(?, ?, ?)",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(equipeId, usuarioId, usuarioNick); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Equipe) Remover(equipeId, usuarioId uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM usuarios_equipe WHERE equipes_id = ? AND usuario_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(equipeId, usuarioId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Equipe) BuscarParticipante(equipeId, usuarioId uint64) (modelos.Equipes, modelos.Usuarios, error) {
	linha, erro := repositorio.db.Query("select equipes_id, usuario_id, usuario_nick from usuarios_equipe where equipes_id = ? and usuario_id = ?",
		equipeId, usuarioId,
	)
	if erro != nil {
		return modelos.Equipes{}, modelos.Usuarios{}, erro
	}
	defer linha.Close()

	var Equipe modelos.Equipes
	var Usuario modelos.Usuarios

	if linha.Next() {
		if erro = linha.Scan(
			&Equipe.Id,
			&Usuario.Id,
			&Usuario.Nick,
		); erro != nil {
			return modelos.Equipes{}, modelos.Usuarios{}, erro
		}
	}

	return Equipe, Usuario, nil
}

func (repositorio Equipe) BuscarParticipantesDaEquipe(equipeId uint64) ([]modelos.Usuarios, error) {
	linhas, erro := repositorio.db.Query("select usuario_id, usuario_nick from usuarios_equipe where equipes_id = ?", equipeId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuarios

	for linhas.Next() {
		var usuario modelos.Usuarios

		if erro = linhas.Scan(
			&usuario.Id,
			&usuario.Nick,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}
