package gin

import (
	"apiecommerce2/categoria"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/v2/bson"
	"log"
)

type CategoriaRequest struct {
	ID string `json:"_id"`
	Nome string `json:"nome" binding:"required"`
}

func AdicionarCategoria(ctx context.Context, serviceCategoria categoria.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var categoriaRequest CategoriaRequest

		if err := c.ShouldBindJSON(&categoriaRequest); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida: " + err.Error()})
			return
		}

		categoriaIn := categoria.Categoria{
			Nome: categoriaRequest.Nome,
		}

		if err := serviceCategoria.Adicionar(ctx, categoriaIn); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
				"erro": err,
			})
			c.Abort()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Requisião processada com sucesso!",
			})
		}
	}
}