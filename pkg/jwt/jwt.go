package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExipreDuration = time.Hour * 2

var JWTSalt = []byte("夏天夏天悄悄过去")

type MyClaims struct {
	UserID   uint64 `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID uint64, username string) (string, error) {
	c := MyClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExipreDuration).Unix(),
			Issuer:    "bluebell",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(JWTSalt)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(t *jwt.Token) (interface{}, error) {
		return JWTSalt, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
