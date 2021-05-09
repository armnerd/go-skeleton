package auth

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GetToken 获取token
func GetToken(user int) string {
	authSecret := []byte(os.Getenv("JWT_SECRET"))
	jwtExpires, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES"))
	expiresTime := time.Duration(jwtExpires)

	type appClaims struct {
		User int `json:"user"`
		jwt.StandardClaims
	}

	claims := appClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiresTime * time.Second).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(authSecret)
	return ss
}
