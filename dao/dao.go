package dao

import (
	"GinNGorm/config"
	"GinNGorm/logger"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	Db  *gorm.DB
	err error
)

func init() {

	//Db, err := gorm.Open("mysql", config.Mysqldb)
	Db, err = gorm.Open(mysql.Open(config.Mysqldb), &gorm.Config{})

	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect error": err.Error()})

	}

	if Db.Error != nil {
		logger.Error(map[string]interface{}{"database error": Db.Error})

	}

	sqlDB, _ := Db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxIdleConns(10)
}
