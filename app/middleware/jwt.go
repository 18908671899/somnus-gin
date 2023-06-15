package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"somnus-gin/app/common/response"
	"somnus-gin/app/services"
	"somnus-gin/global"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			fmt.Println("a")
			response.TokenFail(c)
			c.Abort()
			return
		}
		tokenStr = tokenStr[len(services.TokenType)+1:]

		// Token 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})
		if err != nil {
			fmt.Println(err)
			response.TokenFail(c)
			c.Abort()
			return
		}

		claims := token.Claims.(*services.CustomClaims)
		if claims.Issuer != GuardName {
			fmt.Println("c")
			response.TokenFail(c)
			c.Abort()
			return
		}

		c.Set("token", token)
		c.Set("id", claims.Id)
	}

}
