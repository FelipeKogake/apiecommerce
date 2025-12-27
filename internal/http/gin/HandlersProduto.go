package gin

import (
	produto "apiecommerce2/base_produto"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"log"
	"strconv"
)



func AdicionarProduto(ctx context.Context, s produto.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var produtoRequest produto.Produto

		if err := c.ShouldBindJSON(&produtoRequest); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida: " + err.Error()})
			return
		}

		produtoIn := produto.Produto{
			Nome: produtoRequest.Nome,
			Preco: produtoRequest.Preco,
			Categoria: produtoRequest.Categoria,
		}

		if err := s.Adicionar(ctx, produtoIn); err != nil {
			if err.Error() == "Already Exist" {
				c.JSON(http.StatusOK, gin.H{
					"message": "O produto ja está cadastrado.",
				})
				c.Abort()
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Nao foi possível processar sua requisição.",
				})
				c.Abort()
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Produto adicionado com sucesso!",
			})
		}
	}
}

func ListarProdutos(ctx context.Context, s produto.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var skip int64 = 0
		if page := c.Query("page"); page != "" {
			pageInt, _ := strconv.ParseInt(page, 10, 64)
			calcPage := pageInt*4-4
			skip = calcPage
		}

		lista, err := s.Listar(ctx, skip)
		if err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Nao foi possível processar sua requisição.",
				})
			} else if lista == nil {
				c.JSON(http.StatusOK, gin.H{
					"message": "Não há produtos cadastrados",
				})
			} else {
				c.IndentedJSON(http.StatusOK, lista)
			}
	}
}

func BuscarProduto(ctx context.Context, s produto.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		produto, err := s.Buscar(ctx, id)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else {
			c.JSON(http.StatusOK, produto)
		}
	}
}

func AtualizarProduto(ctx context.Context, s produto.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var produtoUpdate produto.Produto

		idIn := c.Param("id")

		id, err := bson.ObjectIDFromHex(idIn)
		if err != nil {
			log.Println(err)
		}

		produtoUpdate.ID = id

		if err := c.ShouldBindJSON(&produtoUpdate); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida: " + err.Error()})
			return
		}

		if err := s.Atualizar(ctx, produtoUpdate); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Informacoes atualizadas", 
			})
		}
		
	}
}

func DeletarProduto(ctx context.Context, s produto.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := s.Deletar(ctx, id); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Informacoes deletadas!", 
			})
		}
	}
}

func ListarQuantidadesPorCategoria(ctx context.Context, serviceProduto produto.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		if lista, err := serviceProduto.ListarQuantidadesPorCategoria(ctx); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else if lista == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Não há produtos cadastrados",
			})
		} else {
			c.IndentedJSON(http.StatusOK, lista)
		}
	}
}

func BuscarPorNome(ctx context.Context, serviceProduto produto.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		busca := c.Param("busca")

		if lista, err := serviceProduto.BuscarPorNome(ctx, busca); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else {
			c.IndentedJSON(http.StatusOK, lista)
		}
	}
}

func BuscarPorFaixaPreco(ctx context.Context, serviceProduto produto.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var faixaPreco produto.FaixaPreco
		min, _ := strconv.ParseFloat(c.Query("min"), 64)
		faixaPreco.Min = min
		max, _ := strconv.ParseFloat(c.Query("max"), 64)
		faixaPreco.Max = max

		if lista, err := serviceProduto.BuscaPorFaixaPreco(ctx, faixaPreco); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else {
			c.IndentedJSON(http.StatusOK, lista)
		}
	}
}