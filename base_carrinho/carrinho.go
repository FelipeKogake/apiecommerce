package carrinho

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"apiecommerce2/base_produto"
	"apiecommerce2/internal/usuario"
)

type Carrinho struct {
	ID bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Itens []ItemCarrinho `json:"itens" bson:"itens"`
	Usuario usuario.Usuario `json:"usuario" bson:"usuario`
	Total float64 `json:"total" bson:"total"`
}

type ItemCarrinho struct {
	Produto produto.Produto `json:"produto" bson:"produto" binding:"required"`
	Quantidade int `json:"quantidade" bson:"quantidade" binding:"required,gt=0`
}