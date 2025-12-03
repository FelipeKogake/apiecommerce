package produto

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Produto struct {
	ID string `bson:"omitempty"`
	Nome string 
	Preco float64 
	Id_categoria bson.ObjectID 
}