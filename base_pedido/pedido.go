package pedido

import (
	"apiecommerce2/base_carrinho"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Pedido struct {
	ID bson.ObjectID `json:"_id" bson:"_id,omitempty"`
	Data time.Time `json:"data bson:"data"`
	Carrinho carrinho.Carrinho `json:"carrinho" bson:"carrinho" binding:"required"`
	Status string `json:"status" bson:"status"`
	Endereco string `json:"endereco" bson:"endereco`
}

