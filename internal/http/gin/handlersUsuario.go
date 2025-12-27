package gin

import (
	"apiecommerce2/internal/usuario"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CriarUsuario(ctx context.Context, serviceUsuario usuario.UseCase) gin.HandlerFunc{
	return func(c *gin.Context) {
		var usuario usuario.Usuario
		if err := c.ShouldBindBodyWithJSON(&usuario); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida: " + err.Error()})
			return
		}


		if err := serviceUsuario.Adicionar(ctx, usuario); err != nil {
			if err.Error() == "Usuario já existente" {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Esse nome de usuário ja está em uso.",
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Nao foi possível processar sua requisição.",
				})
				c.Abort()
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Usuario criado!",
			})
		}
	}
}

func AtualizarUsuario(ctx context.Context, serviceUsuario usuario.UseCase) gin.HandlerFunc{
	return func(c *gin.Context) {
		var usuario usuario.Usuario
		if err := c.ShouldBindBodyWithJSON(&usuario); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida: " + err.Error()})
			return
		}

		id := LocalizarUsuario(c)
		usuario.ID = id

		if err := serviceUsuario.Atualizar(ctx, usuario); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
			c.Abort()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Informações atualizadas com sucesso!",
			})
		}
	}
}

func DeletarUsuario(ctx context.Context, serviceUsuario usuario.UseCase) gin.HandlerFunc{
	return func(c *gin.Context) {
		id := LocalizarUsuario(c)

		if err := serviceUsuario.Deletar(ctx, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
			c.Abort()
		} else {
			c.SetCookie(
				"session_token", // Nome do cookie (deve corresponder)
				"",              // O valor não importa, pode ser vazio
				-1,              // MaxAge: -1 instrui o navegador a deletar o cookie
				"/",             // Path (deve corresponder)
				"localhost",     // Domain (deve corresponder)
				false,
				true,
			)
			c.JSON(http.StatusOK, gin.H{
				"message": "Usuário deletado",
			})
		}
	}
}

func ListarEnderecos(ctx context.Context, serviceUsuario usuario.UseCase) gin.HandlerFunc{
	return func(c *gin.Context) {
		id := LocalizarUsuario(c)

		if lista, err := serviceUsuario.ListarEnderecos(ctx, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
			c.Abort()
		} else {
			c.IndentedJSON(http.StatusOK, lista)
		}
	}
}

func AdicionarEndereco(ctx context.Context, serviceUsuario usuario.UseCase) gin.HandlerFunc{
	return func(c *gin.Context) {
		var endereco usuario.Endereco
		if err := c.ShouldBindBodyWithJSON(&endereco); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida: " + err.Error()})
			return
		}

		id := LocalizarUsuario(c)

		if err := serviceUsuario.AdicionarEndereco(ctx, id, endereco.Endereco); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
			c.Abort()
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "endereco adicionado",
			})
		}
	}
}

func AtualizarEndereco(ctx context.Context, serviceUsuario usuario.UseCase) gin.HandlerFunc{
	return func(c *gin.Context) {
		id := LocalizarUsuario(c)

		var endereco usuario.Endereco
		if err := c.ShouldBindBodyWithJSON(&endereco); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida: " + err.Error()})
			return
		}

		if err := serviceUsuario.AtualizarEndereco(ctx, id, endereco); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
			c.Abort()
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "Endereco Atualizado",
			})
		}
	}
}

func DeletarEndereco(ctx context.Context, serviceUsuario usuario.UseCase) gin.HandlerFunc{
	return func(c *gin.Context) {
		id := LocalizarUsuario(c)

		indice := c.Param("indice")

		if err := serviceUsuario.RemoverEndereco(ctx, id, indice); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nao foi possível processar sua requisição.",
			})
			c.Abort()
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "Endereco Removido!",
			})
		}
	}
}