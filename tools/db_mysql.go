package tools

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var MYSQLDB *sql.DB

func InitMysql() {
	db, err := sql.Open("mysql", "root:singularity@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	MYSQLDB = db
}
