package main

import (
	"gin-swager-demo/controller"
	_ "gin-swager-demo/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title swageer demo
// @version 3.0
// @description go lang gin web swager demo
// @termsOfService http://127.0.0.1:8080
func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/get_user_info", controller.GetUserInfo)
	r.Run()
}
