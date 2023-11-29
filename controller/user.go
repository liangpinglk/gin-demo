package controller

import (
	"fmt"
	"gin-demo/models"
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
// @Tags User
// @Param Authorization header string true "jwt token"
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
	// user exist
	queryResult := tools.OrmDb.Where("name = ?", userInfo.Name).First(&models.UserInfo{})
	if queryResult.Error == nil {
		tools.HttpJson(c, userInfo, fmt.Sprintf("user %s exist, can't repeat", userInfo.Name), 400)
		return
	}

	// todo: encrypt password
	insertResult := tools.OrmDb.Create(&models.UserInfo{Name: userInfo.Name, Password: userInfo.Password})

	if insertResult.Error != nil {
		tools.HttpJson(c, userInfo, fmt.Sprintf("create failed: %s", insertResult.Error), 400)
		return
	}
	tools.HttpJson(c, userInfo, "create ok", 200)
}

// @Summary get user info
// @Produce json
// @Tags User
// @Param Authorization header string true "jwt token"
// @Param name query string false "用户名"
// @Param page query int false "page"
// @Param page_size query int false "page size"
// @Success 200 {object} openapi.ListUserRes
// @Router /user [get]
func GetUserInfo(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	var totalCount int64
	query := tools.OrmDb.Model(&models.UserInfo{})
	if name != "" {
		query = query.Where("name like ?", fmt.Sprintf("%%%s%%", name))
	}
	query.Count(&totalCount)
	var UserInfoList []models.UserInfo
	query.Limit(pageSize).Offset((page - 1) * pageSize).Find(&UserInfoList)
	result := make(map[string]any)
	result["user_info"] = UserInfoList
	result["count"] = totalCount
	tools.HttpJson(c, result, "get user info successfully", 200)
}

// @Summary update user
// @Param Authorization header string true "jwt token"
// @Tags User
// @Produce json
// @Accept json
// @Param request body openapi.UpdateInfo true "更新用户请求体"
// @Success 200 {object} openapi.UpdateUserRes
// @Router /user [put]
func UpdateUserInfo(c *gin.Context) {
	var updateArg openapi.UpdateInfo
	if err := c.ShouldBindJSON(&updateArg); err != nil {
		tools.HttpJson(c, updateArg, fmt.Sprintf("error arg: %s", err), 400)
		return
	}
	userInfo := models.UserInfo{}
	tools.OrmDb.First(&userInfo, updateArg.ID)

	if updateArg.Name != "" {
		userInfo.Name = updateArg.Name
	}
	if updateArg.Password != "" {
		userInfo.Password = updateArg.Password
	}
	updateResult := tools.OrmDb.Save(&userInfo)

	if updateResult.Error != nil {
		tools.HttpJson(c, updateArg, fmt.Sprintf("error arg: %s", updateResult.Error), 400)
		return
	}
	tools.HttpJson(c, updateArg, fmt.Sprintf("update successfully"), 200)
}

// @Summary login
// @Tags User
// @Produce json
// @Accept json
// @Param request body User true "登陆信息"
// @Success 200 {object} openapi.LoginRes
// @Router /login [post]
func Login(c *gin.Context) {
	var loginInfo User
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		tools.HttpJson(c, loginInfo, fmt.Sprintf("login error: %s", err), 400)
		return
	}
	userInfo := models.UserInfo{}
	queryResult := tools.OrmDb.Where("name = ?", loginInfo.Name).Where("password = ?", loginInfo.Password).First(&userInfo)

	if queryResult.Error != nil {
		tools.HttpJson(c, loginInfo, fmt.Sprintf("login error %s", queryResult.Error), 400)
		return
	}
	claims := tools.MyCustomClaims{
		userInfo.Name, userInfo.ID,
		tools.JWTRegisteredClaims(),
	}
	ss, jwtErr := tools.GenerateJWT(claims)
	if jwtErr != nil {
		tools.HttpJson(c, loginInfo, fmt.Sprintf("login error: %s", jwtErr), 400)
		return
	}
	tools.HttpJson(c, ss, fmt.Sprintf("login successfully"), 200)
}
