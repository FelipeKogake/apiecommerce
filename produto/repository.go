package produto

import (
	"context"
)

type Reader interface {
	Buscar(ctx context.Context, id string) (Produto, error)
	Listar(ctx context.Context) ([]Produto, error)
}

type Writer interface {
	Adicionar(ctx context.Context, produto Produto) error
	Atualizar(ctx context.Context, produto Produto) error
	Deletar(ctx context.Context, id string) error
}

type Repository interface {
	Reader
	Writer
	Close(ctx context.Context) error
}