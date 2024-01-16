package ioc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"passkey-demo/internal/repository/dao"
	"passkey-demo/pkg/confer"
)

func InitDB() *gorm.DB {
	DSN := confer.C().Mysql.DSN
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
