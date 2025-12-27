package avaliacao

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Reader interface {
	Listar(ctx context.Context, filtro mongo.Pipeline) ([]bson.D, error)
}

type Writer interface {
	AdicionarAvaliacao(ctx context.Context, avaliacao Avaliacao) error
}

type Repository interface {
	Reader
	Writer
	Close(ctx context.Context) error
}