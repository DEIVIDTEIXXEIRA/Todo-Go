package modelos

//Usuario representa o formato do usuario da aplicação
type Usuarios struct {
	Id    uint64 `json:"id,omitempty"`
	Nome  string `json:"nome,omitempty"`
	Nick  string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
}
