package usuario

import (
	// "go.mongodb.org/mongo-driver/v2/ bson"
)

type Endereco struct {
	Endereco string `json:"endereco"`
	Indice int `json:"indice"`
}

type EnderecoCarrinho struct {
	Endereco string `json:"endereco" binding:"required"`
}	

