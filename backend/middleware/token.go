package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

var secretKey = []byte("todolist")

type CustomClaims struct {
	Name string `json:"name"` // name
	Id   int64  `json:"id"`   // id
	jwt.RegisteredClaims
}

func GenerateToken(name string, id int64) string {
	log.Printf("Generate token with name:%v and id:%v\n", name, id)
	claims := CustomClaims{
		name,
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err := token.SignedString(secretKey); err != nil {
		log.Printf("Fail to generate token, err: %s\n", err.Error())
		return "fail"
	} else {
		return tokenString
	}
}
