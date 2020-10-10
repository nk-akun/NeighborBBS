package util

import (
	"database/sql"

	"github.com/nk-akun/NeighborBBS/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	db    *gorm.DB
	sqlDB *sql.DB
)

// OpenDB create a connection of db
func OpenDB(dsn string, config *gorm.Config, maxIdleConns, maxOpenConns int, models ...interface{}) (err error) {
	if config == nil {
		config = &gorm.Config{}
	}

	if config.NamingStrategy == nil {
		config.NamingStrategy = schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		}
	}

	if db, err = gorm.Open(mysql.Open(dsn), config); err != nil {
		logs.Logger.Errorf("open database failed: %s", err.Error())
		return
	}

	if sqlDB, err = db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(maxIdleConns)
		sqlDB.SetMaxOpenConns(maxOpenConns)
	} else {
		logs.Logger.Error(err)
	}

	if err = db.AutoMigrate(models...); nil != err {
		logs.Logger.Errorf("auto migrate tables failed: %s", err.Error())
	}
	return
}

// DB return the pointer of db
func DB() *gorm.DB {
	return db
}

// CloseDB close the connection of db
func CloseDB() {
	if sqlDB == nil {
		return
	}
	if err := sqlDB.Close(); nil != err {
		logs.Logger.Errorf("Disconnect from database failed: %s", err.Error())
	}
}
