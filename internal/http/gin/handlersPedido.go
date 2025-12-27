package gin

import (
	"apiecommerce2/base_carrinho"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"log"
	"apiecommerce2/base_pedido"
	// "github.com/golang-jwt/jwt/v5"
	// "os"
	"apiecommerce2/internal/usuario"
)

func AdicionarPedido(ctx context.Context, servicePedido pedido.UseCase, serviceCarrinho carrinho.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var endereco usuario.EnderecoCarrinho
		if err := c.ShouldBindJSON(&endereco); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida: " + err.Error()})
			return
		}

		id := LocalizarUsuario(c)

		total, err10 := serviceCarrinho.ValorTotal(ctx, id)
		if err10 != nil {
			log.Println(err10)
		}

		log.Println(total)

		carrinho, _ := serviceCarrinho.RemoverCarrinho(ctx, id)
	
		valor, err2 := total[0].Value.(float64)
		log.Println(valor)
		if err2 != false {
			log.Println(err2)
		}

		carrinho.Total = valor

		log.Println(carrinho)

		if err := servicePedido.CriarPedido(ctx, carrinho, id, endereco.Endereco); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "Pedido feito!",
			})
		}
	}
}

func ListarPedidos(ctx context.Context, servicePedido pedido.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := LocalizarUsuario(c)

		if lista, err := servicePedido.ListarPedidos(ctx, id); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else if lista == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Não há pedidos ainda.",
			})
		} else {
			c.IndentedJSON(http.StatusOK, lista)
		}

	}
}

func BuscarPedido(ctx context.Context, servicePedido pedido.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		idIn := c.Param("id")

		id, err := bson.ObjectIDFromHex(idIn)
		if err != nil {
			log.Println(err)
		}

		if pedido, err := servicePedido.BuscarPedido(ctx, id); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else {
			c.JSON(http.StatusOK, pedido)
		}
	}
}

func AtualizarStatus(ctx context.Context, servicePedido pedido.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := LocalizarUsuario(c)
		type Status struct {
			ID string `json:"_id"`
			Status string `json:"status`
		}

		var status Status

		if err := c.ShouldBindJSON(&status); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida: " + err.Error()})
			return
		}

		idPed := c.Param("id")

		idPed2, err := bson.ObjectIDFromHex(idPed)
		if err != nil {
			log.Println(err)
		}

		if err := servicePedido.AtualizarStatus(ctx, id, idPed2, status.Status); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "Status atualizado!",
			})
		}
	}
}

func ProdutosMaisVendidos(ctx context.Context, servicePedido pedido.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		if lista, err := servicePedido.ProdutosMaisVendidos(ctx); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else if lista == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Não há produtos cadastrados.",
			})
		} else {
			c.IndentedJSON(http.StatusOK, lista)
		}
	}
}

func VendasPorPeriodo(ctx context.Context, servicePedido pedido.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		if lista, err := servicePedido.VendasPorPeriodo(ctx); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
		} else {
			c.IndentedJSON(http.StatusOK, lista)
		}
	}
}