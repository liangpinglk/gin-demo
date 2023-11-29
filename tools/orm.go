package tools

import (
	"fmt"
	"gin-demo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var OrmDb *gorm.DB

func InitOrmDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", Config["DB_USER"],
		Config["DB_PASSWORD"], Config["DB_HOST"], Config["DB_PORT"], Config["DB_NAME"])
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	OrmDb = db
}

func Migrate() {
	OrmDb.AutoMigrate(&models.UserInfo{})
}
