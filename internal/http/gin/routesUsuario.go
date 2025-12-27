package gin

import (
	"apiecommerce2/internal/usuario/mongodb"
	"context"
	"apiecommerce2/internal/usuario"
	"github.com/gin-gonic/gin"
)

func RoutesUsuario(ctx context.Context, g *gin.Engine, r *mongodb.Repository, serviceUsuario usuario.UseCase) {
	g.POST("/usuario", CriarUsuario(ctx, serviceUsuario))
	g.POST("login", GerarToken(ctx, r))
	g.PUT("/usuario", ValidarToken(ctx), AtualizarUsuario(ctx, serviceUsuario))
	g.DELETE("/usuario", ValidarToken(ctx), DeletarUsuario(ctx, serviceUsuario))
	g.GET("/usuario/endereco", ValidarToken(ctx), ListarEnderecos(ctx, serviceUsuario))
	g.POST("/usuario/endereco", ValidarToken(ctx), AdicionarEndereco(ctx, serviceUsuario))
	g.PUT("/usuario/endereco", ValidarToken(ctx), AtualizarEndereco(ctx, serviceUsuario))
	g.DELETE("/usuario/endereco/:indice", ValidarToken(ctx), DeletarEndereco(ctx, serviceUsuario))
}