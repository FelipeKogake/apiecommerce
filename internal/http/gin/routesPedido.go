package gin

import (
	"context"
	"apiecommerce2/base_pedido"
	"apiecommerce2/base_carrinho"
	"github.com/gin-gonic/gin"
)

func RoutesPedido(ctx context.Context, g *gin.Engine, servicePedido pedido.UseCase, serviceCarrinho carrinho.UseCase) {
	g.POST("/pedido/fazer", ValidarToken(ctx), AdicionarPedido(ctx, servicePedido, serviceCarrinho))
	g.GET("/pedido", ValidarToken(ctx), ListarPedidos(ctx, servicePedido))
	g.GET("/pedido/:id", ValidarToken(ctx), BuscarPedido(ctx, servicePedido))
	g.PUT("/pedido/:id", ValidarToken(ctx), AtualizarStatus(ctx, servicePedido))
	g.GET("/pedido/produtosMaisVendidos", ValidarToken(ctx), ProdutosMaisVendidos(ctx, servicePedido))
	g.GET("/pedido/vendasPorPeriodo", ValidarToken(ctx), VendasPorPeriodo(ctx, servicePedido))
}