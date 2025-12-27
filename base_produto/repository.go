package produto

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Reader interface {
	Buscar(ctx context.Context, id string) (Produto, error)
	Listar(ctx context.Context, filtro mongo.Pipeline) ([]bson.D, error)
}

type Writer interface {
	Adicionar(ctx context.Context, produto Produto) error
	Atualizar(ctx context.Context, produto Produto, filtro bson.M) error
	Deletar(ctx context.Context, filtro bson.M) error
}

type Repository interface {
	Reader
	Writer
	Close(ctx context.Context) error
}