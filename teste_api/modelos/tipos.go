package modelos

type Contatos struct {
	ID      int64  `json:"Id,omitempty"`
	Nome    string `json:"name,omitempty"`
	Celular string `json:"cellphone,omitempty"`
}
