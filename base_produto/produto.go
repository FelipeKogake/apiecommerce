package produto

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	categoria "apiecommerce2/base_categoria"
)

type Produto struct {
	ID bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Nome string `bson:"nome" json:"nome"`
	Preco float64 `bson:"preco" json:"preco" binding:"required,gt=0"`
	Categoria categoria.Categoria `bson:"categoria" json:"categoria"`
}



