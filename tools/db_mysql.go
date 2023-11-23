package tools

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var MYSQLDB *sql.DB

func InitMysql() {
	db, err := sql.Open(Config["DB_DRIVER"], fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", Config["DB_USER"],
		Config["DB_PASSWORD"], Config["DB_HOST"], Config["DB_PORT"], Config["DB_NAME"]))
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	MYSQLDB = db
}
