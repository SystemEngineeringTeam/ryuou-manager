package dboperation

import "github.com/jinzhu/gorm"

const (
	// DB type
	DB_CONNECT = "mysql"

	// DB settings
	// host
	DB_HOST = "localhost"
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

func gormConnect() *gorm.DB {
	db, err := gorm.Open(DB_CONNECT, DB_CONNECTION)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
