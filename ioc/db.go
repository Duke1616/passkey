package ioc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"passkey-demo/config"
	"passkey-demo/internal/repository/dao"
)

func InitDB() *gorm.DB {
	DSN := config.C().Mysql.DSN
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
