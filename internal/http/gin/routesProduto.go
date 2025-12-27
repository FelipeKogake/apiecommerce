package gin

import (
	"context"
	"apiecommerce2/base_produto"
	"github.com/gin-gonic/gin"
)

func RoutesProduto(ctx context.Context, g *gin.Engine, serviceProduto produto.UseCase) {
	g.POST("produto", ValidarToken(ctx), AdicionarProduto(ctx, serviceProduto))
	g.GET("/produto", ValidarToken(ctx), ListarProdutos(ctx, serviceProduto))
	g.GET("/produto/:id", ValidarToken(ctx), BuscarProduto(ctx, serviceProduto))
	g.PUT("/produto/:id", ValidarToken(ctx), AtualizarProduto(ctx, serviceProduto))
	g.DELETE("/produto/:id", ValidarToken(ctx), DeletarProduto(ctx, serviceProduto))
	g.GET("/categoria/quantidadesPorCategoria", ValidarToken(ctx), ListarQuantidadesPorCategoria(ctx, serviceProduto))
	g.GET("/produto/busca/:busca", ValidarToken(ctx), BuscarPorNome(ctx, serviceProduto))
	g.GET("produto/faixaPreco", ValidarToken(ctx), BuscarPorFaixaPreco(ctx, serviceProduto))
}