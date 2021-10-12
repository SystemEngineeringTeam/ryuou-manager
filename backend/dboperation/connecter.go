package dboperation

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	// DB type
	DB_TYPE = "mysql"

	// DB settings
	// host
	DB_HOST = "127.0.0.1"
	// port
	DB_PORT = "33061"
	// username
	DB_USER = "root"
	// password
	DB_PASS = "my-secret-pw"
	// db name
	DB_NAME = "ryuou_manager_db"

	// DB connection settings
	DB_CONNECTION = DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8&parseTime=True&loc=Local"
)

func init() {
	fmt.Println("DB URL is ", DB_CONNECTION)
}

func gormConnect() *gorm.DB {
	db, err := gorm.Open(DB_TYPE, DB_CONNECTION)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
