package controller

import (
	"fmt"
	"gin-demo/tools"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary create user
// @tag user
// @Produce json
// @Accept json
// @Param request body User true "创建用户请求体"
// @Success 200 {object} openapi.CreateUserRes
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var userInfo User
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		tools.HttpJson(c, userInfo, fmt.Sprintf("error arg: %s", err), 400)
		return
	}
	// insert uer info
	// user exist
	querySQL := fmt.Sprintf("select id, name from user where name='%s'", userInfo.Name)
	var name string
	var id int
	err := tools.MYSQLDB.QueryRow(querySQL).Scan(&id, &name)
	if err == nil {
		fmt.Println(id, name)
		tools.HttpJson(c, userInfo, fmt.Sprintf("user %s exist, can't repeat", name), 400)
		return
	}

	// todo: encrypt password
	insertSql := fmt.Sprintf("INSERT INTO user (name, password) VALUES ( '%s', '%s' )", userInfo.Name, userInfo.Password)
	fmt.Println(insertSql)
	_, err = tools.MYSQLDB.Query(insertSql)
	if err != nil {
		tools.HttpJson(c, userInfo, fmt.Sprintf("create failed: %s", err), 400)
		return
	}
	tools.HttpJson(c, userInfo, "create ok", 200)
}
