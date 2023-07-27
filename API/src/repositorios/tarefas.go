package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Tarefas struct {
	db *sql.DB
}

func NovoRepositorioDeTarefas(db *sql.DB) *Tarefas {
	return &Tarefas{db}
}

func (repositorio Tarefas) CriarTarefa(tarefa modelos.Tarefas) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into tarefas (tarefa, observacao, prazo, autor_id) value(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(tarefa.Tarefa, tarefa.Obsevacao, tarefa.Prazo, tarefa.AutorId)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}
