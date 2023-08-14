package jwttoken

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// token := jwt.New(jwt.SigningMethodHS384)
// claims := token.Claims.(jwt.MapClaims)
// claims["email"] = email
// claims["level"] = "admin"
// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

// t, err := token.SignedString([]byte("secret"))

func GenerateJWT(email string) (tokenString string, err error) {
	expirationTime := time.Now().Add(72 * time.Hour)
	claims := &JWTClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
