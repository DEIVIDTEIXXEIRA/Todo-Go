package modelos

import (
	"errors"
	"strings"
)

// Usuario representa o formato do usuario da aplicação
type Usuarios struct {
	Id    uint64 `json:"id,omitempty"`
	Nome  string `json:"nome,omitempty"`
	Nick  string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
}

func (usuario *Usuarios) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuarios) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O campo nome é obrigatório e não pode ficar em branco!")
	}

	if usuario.Nick == "" {
		return errors.New("O campo nick é obrigatório e não pode ficar em branco!")
	}

	if usuario.Email == "" {
		return errors.New("O campo email é obrigatório e não pode ficar em branco!")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("O campo senha é obrigatório e não pode ficar em branco!")
	}

	return nil
}

func (usuario *Usuarios) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
