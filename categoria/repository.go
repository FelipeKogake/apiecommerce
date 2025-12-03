package categoria

import (
	"context"
)

type Reader interface {
	Buscar(ctx context.Context, id string) (Categoria, error)
	Listar(ctx context.Context) ([]Categoria, error)
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