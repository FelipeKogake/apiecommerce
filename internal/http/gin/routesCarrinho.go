package gin

import (
	"context"
	"apiecommerce2/base_carrinho"
	"github.com/gin-gonic/gin"
)

func RoutesCarrinho(ctx context.Context, g *gin.Engine, serviceCarrinho carrinho.UseCase) {
	g.GET("/carrinho", ValidarToken(ctx), ListarItens(ctx, serviceCarrinho))
	g.POST("/carrinho", ValidarToken(ctx), AdicionarItem(ctx, serviceCarrinho))
	g.PUT("/carrinho", ValidarToken(ctx), AtualizarQuantidade(ctx, serviceCarrinho))
	g.DELETE("/carrinho/:indice", ValidarToken(ctx), RemoverItem(ctx, serviceCarrinho))
	g.GET("/carrinho/valorTotal", ValidarToken(ctx), CalcularValorTotal(ctx, serviceCarrinho))
}