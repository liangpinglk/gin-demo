package middleware

import (
	"fmt"
	"gin-demo/tools"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			//c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			//	"error": "Unauthorized",
			//})
			tools.HttpJson(c, nil, fmt.Sprintf("must login: not have token"), 400)
			c.Abort()
			return
		}
		// 验证token
		_, err := tools.ParseJWT(token)
		if err != nil {
			tools.HttpJson(c, nil, fmt.Sprintf("error token: %s", err), 400)
			c.Abort()
			return
		}
		c.Next()
	}
}
