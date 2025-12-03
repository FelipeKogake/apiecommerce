package gin

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"context"

)

type Usuario struct {
	ID string
	Nome string 
	Senha string
}

 
type Claims struct {
	Sub string
	Usuario Usuario
	jwt.RegisteredClaims
}

func GerarToken(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var credenciais Usuario
		json.NewDecoder(c.Request.Body).Decode(&credenciais)

		cred, err := VerificarUsuario(ctx, credenciais);
		if  err != nil {
			return 
		} else if cred.ID == ""{
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"code":400,
				"type":"BadRequest",
				"message":"Credenciais incorretas.",
			})
			c.Abort()
			return 
		}

		expirationTime := time.Now().Add(1 * time.Hour)
		claims := Claims{
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

		c.JSON(http.StatusOK, gin.H{
			"token":tokenString,
		})
	}
}


func ValidarToken(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
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
	
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
    	return []byte(os.Getenv("CHAVE_SECRETA")), nil
	})

	if err != nil || !tkn.Valid {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
		"code":401,
		"type":"Unethorized",
		"message":"Voce nao esta logado.",
	})
		c.Abort()
		return
	}
}
}

func VerificarUsuario(ctx context.Context, usuario Usuario) (Usuario, error) {return Usuario{}, nil}