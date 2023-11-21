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
// @description golang gin web swager demo
// @termsOfService http://127.0.0.1:8080
func main() {
	// init
	tools.InitLog()
	tools.InitMysql()
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(middleware.CatchError(http.StatusInternalServerError))
	r.POST("/login", controller.Login)
	r.Use(middleware.AuthMiddleware())
	r.POST("/user", controller.CreateUser)
	r.GET("/user", controller.GetUserInfo)
	r.PUT("/user", controller.UpdateUserInfo)
	r.Run()
}
