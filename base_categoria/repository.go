package categoria

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	// "go.mongodb.org/mongo-driver/v2/bson"
)

type Reader interface {
	Buscar(ctx context.Context, id string) (Categoria, error)
	Listar(ctx context.Context, filtro mongo.Pipeline) ([]Categoria, error)
}

type Writer interface {
	Adicionar(ctx context.Context, produto Categoria) error
	Atualizar(ctx context.Context, produto Categoria) error
	Deletar(ctx context.Context, id string) error
}

type Repository interface {
	Reader
	Writer
	Close(ctx context.Context) error
}