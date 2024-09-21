package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	AccessTokenExipreDuration  = time.Hour * 24
	RefreshTokenExipreDuration = time.Hour * 24 * 7
)

var JWTSalt = []byte("夏天夏天悄悄过去")

type MyClaims struct {
	UserID   uint64 `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID uint64, username string) (aToken, rToken string, err error) {
	c := MyClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenExipreDuration).Unix(),
			Issuer:    "bluebell",
		},
	}
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(JWTSalt)
	if err != nil {
		return
	}
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(RefreshTokenExipreDuration).Unix(),
		Issuer:    "bluebell",
	}).SignedString(JWTSalt)

	return
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

func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	if _, err = jwt.Parse(rToken, func(t *jwt.Token) (interface{}, error) {
		return JWTSalt, nil
	}); err != nil {
		return
	}
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, func(t *jwt.Token) (interface{}, error) {
		return JWTSalt, nil
	})
	v, _ := err.(*jwt.ValidationError)
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserID, claims.Username)
	}
	return
}
