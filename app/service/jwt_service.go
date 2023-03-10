package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type FboxClaims struct {
	Username string `json:"username"`
	Ukey     string `json:"ukey"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var FboxJWTSecret = []byte("ce5a918b7c0afET3328623")

// placeholder
// GenJwtToken issues signed claims for the current account session.
func GenJwtToken(username string, ukey string) (string, error) {
	c := FboxClaims{
		username,
		ukey,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "ForthBox",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(FboxJWTSecret)
}

// placeholder
func ParseJwtToken(token string) (*FboxClaims, error) {
	tk, err := jwt.ParseWithClaims(token, &FboxClaims{}, func(tk *jwt.Token) (i interface{}, err error) {
		return FboxJWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tk.Claims.(*FboxClaims); ok && tk.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}


