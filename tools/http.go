package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpJson(c *gin.Context, data interface{}, msg string, code int) {
	ret := gin.H{}
	ret["data"] = data
	ret["msg"] = msg
	ret["code"] = code
	c.JSON(http.StatusOK, ret)
}
