package gin

import (
	categoria "apiecommerce2/base_categoria"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/v2/bson"
	"log"
)

func AdicionarCategoria(ctx context.Context, serviceCategoria categoria.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var categoriaRequest categoria.CategoriaRequest

		if err := c.ShouldBindJSON(&categoriaRequest); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida: " + err.Error()})
			return
		}

		categoria := categoriaRequest.ToCategoria()

		if err := serviceCategoria.Adicionar(ctx, categoria); err != nil {
			if err.Error() == "Already exists" {
				c.JSON(http.StatusOK, gin.H{
					"message": "A categoria já está cadastrada.",
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Nao foi possível processar sua requisição.",
				})
			} 
			c.Abort()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Categoria Adicionada com sucesso!",
			})
		}
	}
}

func ListarCategorias(ctx context.Context, serviceCategoria categoria.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		lista, err := serviceCategoria.Listar(ctx)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else if lista == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Não categorias criadas.",
			})
		} else {
			c.IndentedJSON(http.StatusOK, lista)
		}
		
	}
}