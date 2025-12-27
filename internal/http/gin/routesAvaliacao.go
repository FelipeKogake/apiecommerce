package gin

import (
	"context"
	"apiecommerce2/base_avaliacao"
	"apiecommerce2/base_produto"
	"github.com/gin-gonic/gin"
)

func RoutesAvaliacao(ctx context.Context, g *gin.Engine, serviceAvaliacao avaliacao.UseCase, serviceProduto produto.UseCase) {
	g.POST("/avaliacao/:idProduto", ValidarToken(ctx), AdicionarAvaliacao(ctx, serviceAvaliacao, serviceProduto))
	g.GET("/avaliacao/:idProduto", ValidarToken(ctx), ListarAvaliacoes(ctx, serviceAvaliacao))
	g.GET("/avaliacao/mediaTempoReal", ValidarToken(ctx), CalcularMediaTempoReal(ctx, serviceAvaliacao))
}