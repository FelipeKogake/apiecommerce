package produto

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	categoria "apiecommerce2/base_categoria"
)

type ProdutoRequest struct {
	ID bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Nome string `bson:"nome" json:"nome"`
	Preco float64 `bson:"preco" json:"preco"`
	Categoria categoria.CategoriaRequest `bson:"categoria" json:"categoria"`
}

type FaixaPreco struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}