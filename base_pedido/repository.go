package pedido

import (
	"apiecommerce2/base_carrinho"
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Reader interface {
	Listar(ctx context.Context, pipeline mongo.Pipeline) ([]bson.D, error)
	BuscarPedido(ctx context.Context, filtro bson.M)  (Pedido, error)
}

type Writer interface {
	CriarPedido(ctx context.Context, carrinho carrinho.Carrinho, pedido Pedido) error 
	AtualizarStatus(ctx context.Context, filtro bson.M, update bson.M) error
}

type Repository interface {
	Reader
	Writer
	Close(ctx context.Context) error
}