package gin

import (
	"apiecommerce2/base_carrinho"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListarItens(ctx context.Context, serviceCarrinho carrinho.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		id := LocalizarUsuario(c)

		itens, err3 := serviceCarrinho.ListarItens(ctx, id)
		if err3 != nil {
			log.Println(err3)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else if itens == nil{
			c.JSON(http.StatusOK, gin.H{
				"message": "O carrinho está vazio, adicione um item para visualizá-lo.",
			})
		} else {
			c.IndentedJSON(http.StatusOK, itens)
		}


	}
}

func AdicionarItem(ctx context.Context, serviceCarrinho carrinho.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {

		var item carrinho.ItemCarrinho
		if err := c.ShouldBindJSON(&item); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida: " + err.Error()})
			return
		}

		itemFinal := carrinho.ItemCarrinho{
			Produto: item.Produto,
			Quantidade: item.Quantidade,
		}

		id := LocalizarUsuario(c)

		err3 := serviceCarrinho.AdicionarItem(ctx, itemFinal, id)
		if err3 != nil {
			log.Println(err3)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "Produto adicionado!",
			})
		}
	}
}

func AtualizarQuantidade(ctx context.Context, serviceCarrinho carrinho.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {

		log.Println("Atualizando item")

		type resposta struct {
			Indice int `json:"indice`
			Quantidade int `json:"quantidade"`
		}

		id := LocalizarUsuario(c)

		var info resposta
		if err := c.ShouldBindJSON(&info); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida: " + err.Error()})
			return
		}

		err5 := serviceCarrinho.AtualizarQuantidade(ctx, info.Indice, info.Quantidade, id)
		if err5 != nil {
			log.Println(err5)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "Quantidade atualizada!",
			})
		}
	}
}

func RemoverItem(ctx context.Context, serviceCarrinho carrinho.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		indice := c.Param("indice")
		indiceInt, _:= strconv.Atoi(indice)

		id := LocalizarUsuario(c)

		if err := serviceCarrinho.RemoverItem(ctx, indiceInt, id); err != nil {
			if err.Error() == "O produto não foi encontrado." {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Nso existe esse produto no carrinho.",
				})
			} else {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{
						"message": "Nao foi possível processar sua requisição.",
					})
			}
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "Item removido!",
			})
		}
	}
}

func CalcularValorTotal(ctx context.Context, serviceCarrinho carrinho.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := LocalizarUsuario(c)

		if valor, err := serviceCarrinho.ValorTotal(ctx, id); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else {
			c.JSON(http.StatusOK, valor)
		}
	}
}