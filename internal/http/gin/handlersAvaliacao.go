package gin

import (
	"apiecommerce2/base_avaliacao"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"log"
	"apiecommerce2/base_produto"
	// "github.com/golang-jwt/jwt/v5"
	// "os"
	// "strconv"
	"apiecommerce2/base_categoria"
)

// Adicionar Avaliacao godoc
// @Summary Adiciona uma nova avaliacao
// @Description Adiciona uma avaliacao especificamente em um produto já existente
// @Tags avaliacao
// @Produce json
// Success 201 {array} error
// @Router /avaliacao [post]
func AdicionarAvaliacao(ctx context.Context, serviceAvaliacao avaliacao.UseCase, serviceProduto produto.UseCase) gin.HandlerFunc{
	return func(c *gin.Context) {
		var avaliacao avaliacao.Avaliacao

		id := LocalizarUsuario(c)

		if err := c.BindJSON(&avaliacao); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida: " + err.Error()})
			return
		}

		produtoID := c.Param("idProduto")
		produtoB, err := serviceProduto.Buscar(ctx, produtoID)
		if err != nil {
			log.Println(err)
		}

		categoriaDTO := categoria.CategoriaRequest{
			Nome: produtoB.Categoria.Nome,
		}

		produtoR := produto.ProdutoRequest{
			ID: produtoB.ID,
			Nome: produtoB.Nome,
			Preco: produtoB.Preco,
			Categoria: categoriaDTO,
		}


		avaliacao.Produto = produtoR
		avaliacao.UsuarioID = id

		if err := serviceAvaliacao.AdicionarAvaliacao(ctx, avaliacao); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "Obrigado pela avaliação!",
			})
		}
	}
}

func ListarAvaliacoes(ctx context.Context, serviceAvaliacao avaliacao.UseCase) gin.HandlerFunc{
	return func(c *gin.Context) {
		idProduto := c.Param("idProduto")

		id, err := bson.ObjectIDFromHex(idProduto)
		if err != nil {
			log.Println(err)
		}

		if lista, err := serviceAvaliacao.ListarAvalicoesProduto(ctx, id); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else if lista == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Não há avaliações para esse produto.",
			})
		}else {
			c.IndentedJSON(http.StatusOK, lista)
		}
	}
}

func CalcularMediaTempoReal(ctx context.Context, serviceAvaliacao avaliacao.UseCase) gin.HandlerFunc{
	return func(c *gin.Context) {
		if lista, err := serviceAvaliacao.MediaTempoReal(ctx); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else {
			c.IndentedJSON(http.StatusOK, lista)
		}
	}
}