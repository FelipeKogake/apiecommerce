package avaliacao

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"apiecommerce2/base_produto"
)

type Avaliacao struct {
	ID bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Produto produto.ProdutoRequest `bson:"produto" json:"produto"`
	Pontos int `bson:"pontos" json:"pontos" binding:"required,gte=0,lte=10"`
	Comentarios string `bson:"comentarios" json:"comentarios" binding:"min=10"`
	UsuarioID bson.ObjectID `bson:"usuarioid" json:"usuarioid"`
}

