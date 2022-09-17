package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"math/rand"
	"taveler/infrastructure/model"
	"time"
)

func RandomUID() string {
	rand.Seed(time.Now().UnixNano())
	min := 1000000000
	max := 9999999999
	return fmt.Sprintf("UID%v", rand.Intn(max-min+1)+min)
}

func RandomCode() string {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	return fmt.Sprintf("%v", rand.Intn(max-min+1)+min)
}

const ttl = time.Hour

func GenerateJWT(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    user.ID,
		"user_role":  user.Role,
		"exp":        time.Now().Add(ttl).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString([]byte("secret"))
}

//func CheckJWT(token string) (*jwt.Auth, error) {
//
//}
