package modelos

type Equipes struct {
	Id        uint64 `json:"id,omitempty"`
	Nome      string `json:"nome,omitempty"`
	Descricao string `json:"descricao,omitempty"`
	AutorId   uint64 `json:"autorId,omitempty"`
}
