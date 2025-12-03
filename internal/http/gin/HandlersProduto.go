package gin

import (
	"apiecommerce2/produto"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"log"
)

type ProdutoRequest struct {
	ID string `json:"_id"`
	Nome string `json:"nome" binding:"required"`
	Preco float64 `json:"preco" binding:"required,gt=0"`
	Id_categoria string `json:"id_categoria" binding:"required"`
}

func AdicionarProduto(ctx context.Context, s produto.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var produtoRequest ProdutoRequest

		if err := c.ShouldBindJSON(&produtoRequest); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida: " + err.Error()})
			return
		}

		id, err := bson.ObjectIDFromHex(produtoRequest.Id_categoria)
		if err != nil {
			log.Println("tamo aqui", err)
		}

		produtoIn := produto.Produto{
			ID: produtoRequest.ID,
			Nome: produtoRequest.Nome,
			Preco: produtoRequest.Preco,
			Id_categoria: id,
		}

		if err := s.Adicionar(ctx, produtoIn); err != nil {
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