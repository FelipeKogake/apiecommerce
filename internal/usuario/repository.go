package usuario

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Verification interface {
	ValidarUsuario(ctx context.Context, usuario Usuario) (Usuario, error)
}

type Reader interface {
	BuscarUsuarioPorNome(ctx context.Context, nome string) (Usuario, error)
	Listar(ctx context.Context, pipeline mongo.Pipeline) ([]bson.D, error)
}

type Write interface {
	CriarUsuario(ctx context.Context, usuario Usuario) error
	Atualizar(ctx context.Context, filtro bson.D, update bson.D) error 
	DeletarUsuario(ctx context.Context, filtro bson.M) error
}

type Repository interface {
	Verification
	Reader
	Write
	Close(ctx context.Context) error
}