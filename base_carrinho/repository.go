package carrinho

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Reader interface {
	ListarItens(ctx context.Context, pipeline mongo.Pipeline) ([]ItemCarrinho, error)
}

type Writer interface {
	CriarCarrinho(ctx context.Context, carrinho Carrinho) error
	RemoverCarrinho(ctx context.Context, filtro bson.M) (Carrinho, error)
	Listar(ctx context.Context, pipeline mongo.Pipeline) ([]bson.D, error)
	Atualizar(ctx context.Context, filtro bson.M, update bson.M) error
}

type Repository interface {
	Reader
	Writer
	Close(ctx context.Context) error
}

