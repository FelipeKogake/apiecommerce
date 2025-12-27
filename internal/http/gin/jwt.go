package gin

import (
	"net/http"
	"time"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"context"
	"apiecommerce2/internal/usuario/mongodb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"apiecommerce2/internal/usuario"
	"log"
)


func GerarToken(ctx context.Context, s *mongodb.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var credenciais usuario.Usuario
		c.BindJSON(&credenciais)

		cred, err := s.ValidarUsuario(ctx, credenciais);
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"code":400,
				"type":"BadRequest",
				"message":"Credenciais incorretas.",
			})
			c.Abort()
			return 
		}

		if  err != nil {
			return 
		} else if cred.ID.IsZero() {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"code":400,
				"type":"BadRequest",
				"message":"Credenciais incorretas.",
			})
			c.Abort()
			return 
		}

		expirationTime := time.Now().Add(1 * time.Hour)

		claims := usuario.Claims{
			Sub: cred.ID,
			Usuario: cred,
			RegisteredClaims:jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}

		godotenv.Load("../resources/.env")	

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(os.Getenv("CHAVE_SECRETA")))

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"code":500,
			"type":"InternalServerError",
			"message":"Não foi possível processar sua solicitação no momento.",
		})
			c.Abort()
			return 
		}

		c.SetCookie(
			"session_token", // Nome do cookie
			tokenString,    // Valor do cookie
			3600,            // maxAge em segundos (1 hora)
			"/",             // Path (disponível em todo o site)
			"localhost",     // Domain (para qual domínio é válido)
			true,            // Secure (enviar apenas sobre HTTPS)
			true,            // HttpOnly (não acessível via JavaScript)
		)

		c.JSON(http.StatusOK, gin.H{
			"message": "Você esta logado!!",
		})
	}
}


func ValidarToken(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
	tokenString, err := c.Cookie("session_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
    return
	}
	if tokenString == "" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
		"code":401,
		"type":"Unethorized",
		"message":"Voce nao esta logado.",
	})
		c.Abort()
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
	}
	
	claims := &usuario.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
    	return []byte(os.Getenv("CHAVE_SECRETA")), nil
	})

	if err != nil || !tkn.Valid {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
		"code":401,
		"type":"Unethorized",
		"message":"Token inválido ou login não afetuado.",
	})
		c.Abort()
	}
	}
}

func LocalizarUsuario(c *gin.Context) bson.ObjectID {
	tokenString, err := c.Cookie("session_token")

		claims := &usuario.Claims{}

		tkn, err2 := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("CHAVE_SECRETA")), nil
		})
		if err2 != nil {
			log.Println(err2)
		}

		if err != nil || !tkn.Valid {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"code":401,
			"type":"Unethorized",
			"message":"Token inválido.",
			})
				c.Abort()
		}
		return claims.Sub
}