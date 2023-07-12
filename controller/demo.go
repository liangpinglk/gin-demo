package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary get user info by name
// @Produce json
// @Param name query string true "用户名"
// @Success 200 {object} map[string]string
// @Router /get_user_info [get]
func GetUserInfo(c *gin.Context) {
	name := c.DefaultQuery("name", "xx")
	UserInfo := map[string]map[string]string{"liangping": {"birthday": "19960702", "sex": "man", "job": "programmer"}, "lianglele": {"birthdat": "19950203", "sex": "man", "job": "civil servant"}}
	result, ok := UserInfo[name]
	if ok {
		c.JSON(http.StatusOK, result)
		return
	}
	c.String(http.StatusOK, "%s not exist", name)
}
