package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeusuairos cria um novo repositorio de usuarios
func NovoRepositorioDeusuairos(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuarios) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Usuarios) BuscarPorId(usuarioId uint64) (modelos.Usuarios, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email from usuarios where id = ?",
		usuarioId,
	)
	if erro != nil {
		return modelos.Usuarios{}, erro 
	}
	defer linhas.Close()

	var usuario modelos.Usuarios

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
		); erro != nil {
			return modelos.Usuarios{}, erro 
		}
	}

	return usuario, nil 

}

func (repositorio Usuarios) Atualizar(Id uint64, usuario modelos.Usuarios) error {
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return  erro 
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, Id); erro != nil {
		return erro 
	}
	
	return nil 
}