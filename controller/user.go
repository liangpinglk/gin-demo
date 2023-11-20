package controller

import (
	"fmt"
	"gin-demo/openapi"
	"gin-demo/tools"
	"github.com/gin-gonic/gin"
	"strconv"
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

// @Summary get user info
// @Produce json
// @Param name query string false "用户名"
// @Param page query int false "page"
// @Param page_size query int false "page size"
// @Success 200 {object} openapi.ListUserRes
// @Router /user [get]
func GetUserInfo(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	querySQL := "select id, name from user where 1=1 "
	if name != "" {
		querySQL += fmt.Sprintf(" and name like \"%%s%\"", name)
	}
	querySQL += fmt.Sprintf("limit %d offset %d ", pageSize, pageSize*(page-1))
	var UserInfoList []any
	results, err := tools.MYSQLDB.Query(querySQL)
	if err != nil {
		tools.HttpJson(c, querySQL, fmt.Sprintf("get user info %s failed", err), 400)
		return
	}

	for results.Next() {
		var user openapi.ListUserInfo
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		UserInfoList = append(UserInfoList, user)
	}

	totalCountSQL := "select count(id) as count from user"
	var count int
	tools.MYSQLDB.QueryRow(totalCountSQL).Scan(&count)
	result := make(map[string]any)
	result["user_info"] = UserInfoList
	result["count"] = count
	tools.HttpJson(c, result, "get user info successfully", 200)
}
