package usuario

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"github.com/golang-jwt/jwt/v5"
)

type Usuario struct {
	ID bson.ObjectID `bson:"_id,omitempty"`
	Nome string `bson:"nome" json:"nome" binding:"required"`
	Senha string `bson:"senha" json:"senha" binding:"required"`
	Enderecos []string `bson:"enderecos" json:"enderecos"`
}

 
type Claims struct {
	Sub bson.ObjectID `bson:"sub" json:"sub"`
	Usuario Usuario `bson:"usuario" json:"usuario"`
	jwt.RegisteredClaims 
}