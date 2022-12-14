package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWTSecret = []byte("!!SECRET!!")

func GenerateJWT(id uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 300).Unix()
	t, _ := token.SignedString(JWTSecret)

	fmt.Println("claims")
	fmt.Println(claims)
	return t
}
