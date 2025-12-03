package usuario

import (
	// "go.mongodb.org/mongo-driver/v2/bson"
	"github.com/golang-jwt/jwt/v5"
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