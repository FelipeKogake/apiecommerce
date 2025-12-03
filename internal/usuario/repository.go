package usuario

import (
	"context"
)

type Verification interface {
	VerificarUsuario(ctx context.Context, usuario Usuario) (Usuario, error)
}

type Write interface {
	CriarUsuario(ctx context.Context, usuario Usuario) error
	Atualizar(ctx context.Context, usuario Usuario) error 
	DeletarUsuario(ctx context.Context, id string) error
}

type Repository interface {
	Verification
	Write
	Close(ctx context.Context) error
}