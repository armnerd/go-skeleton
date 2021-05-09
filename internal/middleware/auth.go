package middleware

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type requestHeader struct {
	Authorization string `header:"Authorization"`
}

// AuthRequired 登录验证
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取header信息
		header := requestHeader{}
		c.BindHeader(&header)
		if len(header.Authorization) == 0 {
			c.JSON(200, gin.H{
				"code":    4003,
				"message": "auth fail: no Authorization header",
				"data":    "",
			})
			c.Abort()
			return
		}

		// 获取secret
		authSecret := []byte(os.Getenv("JWT_SECRET"))

		// 获取jwt负载,验证
		token, err := jwt.Parse(header.Authorization, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return authSecret, nil
		})
		if err != nil {
			c.JSON(200, gin.H{
				"code":    4003,
				"message": fmt.Sprintf("auth fail: %v", err),
				"data":    "",
			})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["user"] == nil {
				c.JSON(200, gin.H{
					"code":    4003,
					"message": "auth fail: no user in claims",
					"data":    "",
				})
				c.Abort()
				return
			}
			c.Set("user", claims["user"])
		} else {
			c.JSON(200, gin.H{
				"code":    4003,
				"message": "auth fail: can not get claims",
				"data":    "",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
