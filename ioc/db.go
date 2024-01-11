package ioc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"passkey-demo/internal/repository/dao"
)

func InitDB() *gorm.DB {
	DSN := "root:ebondorthanc123@tcp(10.31.0.15:9999)/passkey"
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
