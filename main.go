package main

import (
	"gin-demo/controller"
	_ "gin-demo/docs"
	"gin-demo/middleware"
	"gin-demo/tools"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @title swageer demo
// @version 3.0
// @description go lang gin web swager demo
// @termsOfService http://127.0.0.1:8080
func main() {
	// init
	tools.InitLog()
	tools.InitMysql()
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(middleware.CatchError(http.StatusInternalServerError))
	r.GET("/get_user_info", controller.GetUserInfo)
	r.Run()
}
