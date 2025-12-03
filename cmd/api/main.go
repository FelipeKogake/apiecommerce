package main

import (
	"apiecommerce2/config"
	"context"
	"log"
	"apiecommerce2/produto"
	mongodbProduto "apiecommerce2/produto/mongodb"
	internal "apiecommerce2/internal/http/gin"
	"apiecommerce2/categoria"
	mongodbCategoria"apiecommerce2/categoria/mongodb"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Println(err)
	}

	ctx := context.Background()
	repoProduto, _ := mongodbProduto.NewRepository(cfg.Database.Nome, "produto", cfg.Database.URI)
	repoCategoria, _ := mongodbCategoria.NewRepository(cfg.Database.Nome, "categoria", cfg.Database.URI)
	
	defer repoProduto.Close(ctx)
	defer repoCategoria.Close(ctx)
	
	serviceProduto := produto.NewService(repoProduto)
	serviceCategoria := categoria.NewService(repoCategoria)

	g := gin.Default()

	internal.RoutesProduto(ctx, g, serviceProduto)
	internal.RoutesCategoria(ctx, g, serviceCategoria)
	

	
	g.Run(cfg.Port)
}