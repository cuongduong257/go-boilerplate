package driver

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// MysqlDB : Define MysqlDB
type MysqlDB struct {
	Database *gorm.DB
}

// Mysql : Define Mysql
var Mysql = &MysqlDB{}

// Connect : Connect to mysql db
func Connect(host, port, user, password, dbName string) *MysqlDB {
	connStr := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName)
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	Mysql.Database = db
	return Mysql
}
