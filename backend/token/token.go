package token

import (
	"fmt"
	"github.com/DeathVenom54/doto-backend/db"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

type AuthClaims struct {
	jwt.RegisteredClaims
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func ParseToken(token string) (*AuthClaims, error) {
	var claims AuthClaims
	data, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("JWT_KEY"), nil
	})
	if err != nil {
		return nil, err
	}
	if !data.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return &claims, nil
}

func VerifyToken(token string) bool {
	claims, err := ParseToken(token)
	if err != nil {
		return false
	}

	_, err = db.GetUserById(claims.ID)
	if err != nil {
		return false
	}

	return true
}

func CreateToken(claims *AuthClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	return tokenString, err
}
