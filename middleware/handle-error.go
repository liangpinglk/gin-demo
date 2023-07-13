package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CatchError(MsgCode int) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"error": r,
					"code":  MsgCode,
					"data":  gin.H{},
				})
			}
		}()
		c.Next()
	}

}
