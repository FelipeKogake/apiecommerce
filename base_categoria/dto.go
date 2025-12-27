package categoria

import ()

type CategoriaRequest struct {
	Nome string `json:"nome"`
}

type QuantidadeProdutoPorCategoria struct {
	Categoria string `bson:"_id" json:"categoria"`
	Quantidade int `bson:"quantidade"`
}

func (c *CategoriaRequest) ToCategoria() Categoria {
	return Categoria{
		Nome: c.Nome,
	}
}