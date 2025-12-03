package gin

import (
	"context"
	"apiecommerce2/produto"
	"github.com/gin-gonic/gin"
)

func RoutesProduto(ctx context.Context, g *gin.Engine, serviceProduto produto.UseCase) {

	g.POST("produto", AdicionarProduto(ctx, serviceProduto))

}