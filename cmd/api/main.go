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
	mongodbUsuario "apiecommerce2/internal/usuario/mongodb"
	"apiecommerce2/internal/usuario"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Println(err)
	}

	ctx := context.Background()
	repoProduto, _ := mongodbProduto.NewRepository(cfg.Database.Nome, "produto", cfg.Database.URI)
	repoCategoria, _ := mongodbCategoria.NewRepository(cfg.Database.Nome, "categoria", cfg.Database.URI)
	repoUsuario, _ := mongodbUsuario.NewRepository(cfg.Database.Nome, "usuario", cfg.Database.URI)
	
	defer repoProduto.Close(ctx)
	defer repoCategoria.Close(ctx)
	defer repoUsuario.Close(ctx)
	
	serviceProduto := produto.NewService(repoProduto)
	serviceCategoria := categoria.NewService(repoCategoria)
	serviceUsuario := usuario.NewService(repoUsuario)

	g := gin.Default()

	internal.RoutesProduto(ctx, g, serviceProduto)
	internal.RoutesCategoria(ctx, g, serviceCategoria)
	internal.RoutesUsuario(ctx, g, repoUsuario, serviceUsuario)
	

	
	g.Run(cfg.Port)
}