package categoria

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Categoria struct {
	ID bson.ObjectID `bson:"omitempty"`
	Nome string
}