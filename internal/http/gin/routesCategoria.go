package gin

import (
	"context"
	"apiecommerce2/categoria"
	"github.com/gin-gonic/gin"
)

func RoutesCategoria(ctx context.Context, g *gin.Engine, serviceCategoria categoria.UseCase) {

	g.POST("categoria", ValidarToken(ctx), AdicionarCategoria(ctx, serviceCategoria))

}