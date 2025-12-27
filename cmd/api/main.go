package main

import (
	"apiecommerce2/config"
	"context"
	"log"
	"apiecommerce2/base_produto"
	mongodbProduto "apiecommerce2/base_produto/mongodb"
	internal "apiecommerce2/internal/http/gin"
	"apiecommerce2/base_categoria"
	mongodbCategoria"apiecommerce2/base_categoria/mongodb"
	"github.com/gin-gonic/gin"
	mongodbUsuario "apiecommerce2/internal/usuario/mongodb"
	"apiecommerce2/internal/usuario"
	mongoCarrinho "apiecommerce2/base_carrinho/mongodb"
	"apiecommerce2/base_carrinho"
	mongoPedido "apiecommerce2/base_pedido/mongodb"
	"apiecommerce2/base_pedido"
	mongoAvaliacao "apiecommerce2/base_avaliacao/mongodb"
	"apiecommerce2/base_avaliacao"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "apiecommerce2/docs"
)

// @title API de ecommerce
// @version 1.0
// @description Simulacao real dos processos envolvente de uma experiencia completa dentro de uma navegacao de comercio digital
// @contact.name API suport
// @contact.email felipe.kogake@gmail.com
// @servers http://localhost:{port} Servidor de Desenvolvimento Local port=8080 Porta padrão para desenvolvimento

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Println(err)
	}

	ctx := context.Background()
	repoProduto, _ := mongodbProduto.NewRepository(cfg.Database.Nome, "produto", cfg.Database.URI)
	repoCategoria, _ := mongodbCategoria.NewRepository(cfg.Database.Nome, "categoria", cfg.Database.URI)
	repoUsuario, _ := mongodbUsuario.NewRepository(cfg.Database.Nome, "usuario", cfg.Database.URI)
	repoCarrinho, _ := mongoCarrinho.NewRepository(cfg.Database.Nome, "carrinho", cfg.Database.URI)
	repoPedido, _ := mongoPedido.NewRepository(cfg.Database.Nome, "pedido", cfg.Database.URI)
	repoAvaliacao, _ := mongoAvaliacao.NewRepository(cfg.Database.Nome, "avaliacao", cfg.Database.URI)
	
	defer repoProduto.Close(ctx)
	defer repoCategoria.Close(ctx)
	defer repoUsuario.Close(ctx)
	defer repoCarrinho.Close(ctx)
	defer repoPedido.Close(ctx)
	defer repoAvaliacao.Close(ctx)
	
	serviceProduto := produto.NewService(repoProduto)
	serviceCategoria := categoria.NewService(repoCategoria)
	serviceUsuario := usuario.NewService(repoUsuario)
	serviceCarrinho := carrinho.NewService(repoCarrinho)
	servicePedido := pedido.NewService(repoPedido)
	serviceAvaliacao := avaliacao.NewService(repoAvaliacao)

	g := gin.Default()

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	internal.RoutesProduto(ctx, g, serviceProduto)
	internal.RoutesCategoria(ctx, g, serviceCategoria)
	internal.RoutesUsuario(ctx, g, repoUsuario, serviceUsuario)
	internal.RoutesCarrinho(ctx, g, serviceCarrinho)
	internal.RoutesPedido(ctx, g, servicePedido, serviceCarrinho)
	internal.RoutesAvaliacao(ctx, g, serviceAvaliacao, serviceProduto)
	
	g.Run(cfg.Port)
}