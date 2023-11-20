package controller

import (
	"gin-demo/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary get user info by name
// @Produce json
// @Param name query string true "用户名"
// @Success 200 {object} models.UserInfo
// @Router /get_user_info [get]
func getUserInfo(c *gin.Context) {
	//name := c.DefaultQuery("name", "xx")

	panic("force stop ")
	tools.Sugar.Info("get user info log")
	// go-sql-mysql
	//insertSql := fmt.Sprintf("INSERT INTO test1 (`name`) VALUES ('%s');", tools.RandomString(5))
	//fmt.Println(tools.MYSQLDB)
	//insert, err := tools.MYSQLDB.Query(insertSql)
	//if err != nil {
	//	fmt.Println("insert error: ", err)
	//} else {
	//	fmt.Println("insert successfully:", insert)
	//}
	name := c.Query("name")
	userInfo := map[string]map[string]string{"liangping": {"birthday": "199607021", "sex": "man", "job": "programmer"}, "lianglele": {"birthdat": "19950203", "sex": "man", "job": "civil servant"}}
	result, ok := userInfo[name]
	if ok {
		c.JSON(http.StatusOK, result)
		return
	}
	//panic("user not exist") // 触发panic
	c.String(http.StatusOK, "%s not exist", name)
}
