package feiralivre

import "errors"

type FeiraLivre struct {
	ID         int    `json:"id,omitempty"`
	Nome       string `json:"nome"`
	Lat        string `json:"lat"`
	Lng        string `json:"lng"`
	Setcens    string `json:"setcens"`
	Areap      string `json:"areap"`
	Coddist    string `json:"coddist"`
	Distrito   string `json:"distrito"`
	Codsubpref string `json:"codsubpref"`
	Subprefe   string `json:"subprefe"`
	Regiao5    string `json:"regiao5"`
	Regiao8    string `json:"regiao8"`
	Registro   string `json:"registro"`
	Logradouro string `json:"logradouro"`
	Numero     string `json:"numero"`
	Bairro     string `json:"bairro"`
	Referencia string `json:"referencia"`
}

func (f FeiraLivre) Validate() error {

	if f.Nome == "" {
		return errors.New("Nome é obrigatório.")
	}

	return nil
}
