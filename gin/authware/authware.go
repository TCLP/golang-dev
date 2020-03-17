package authware

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	TokenExpireDuration = time.Hour * 2
	TokenSecret         = []byte("123456")
)

type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenToken(username string) (string, error) {
	c := CustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "ylp",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(TokenSecret)
}

func ParseToken(tokenstring string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenstring, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return TokenSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token invalid")
}
