package ioc

import (
	"fmt"
	"github.com/Duke1616/passkey/config"
	"github.com/Duke1616/passkey/internal/repository/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	host := config.C().Mysql.HOST
	port := config.C().Mysql.PORT
	username := config.C().Mysql.USERNAME
	password := config.C().Mysql.PASSWORD
	database := config.C().Mysql.DATABASE
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(DSN))
	if err != nil {
		panic(err)
	}

	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}
