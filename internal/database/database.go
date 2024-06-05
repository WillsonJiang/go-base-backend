package database

import (
	"fmt"
	"time"

	"backend/pkg/env"
	"backend/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var db *gorm.DB

func Init() {
	var err error
	address := fmt.Sprintf(
		"%v:%v@(%v:%v)/%v?charset=%v&loc=Local&parseTime=True",
		env.Get("DATABASE_USERNAME", "user"),
		env.Get("DATABASE_PASSWORD", ""),
		env.Get("DATABASE_HOST", "127.0.0.1"),
		env.Get("DATABASE_PORT", "3306"),
		env.Get("DATABASE_NAME", "wutu"),
		env.Get("DATABASE_CHARSET", "utf8"),
	)

	db, err = gorm.Open(mysql.Open(address), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	if err != nil || db.Error != nil {
		logger.Panic("Connect database faild")
	}

	sqlDB, err := db.DB()
	if err != nil || db.Error != nil {
		logger.Panic("Connect database faild")
	}

	sqlDB.SetMaxOpenConns(env.GetInt("DATABASE_MAX_CONNECTIONS", 0))
	sqlDB.SetMaxIdleConns(env.GetInt("DATABASE_MAX_IDLE_CONNECTIONS", 0))
	sqlDB.SetConnMaxLifetime(
		time.Duration(env.GetInt64("DATABASE_MAX_CONNECTIONS_LIFETIME", 5)) * time.Minute)

	err = db.Use(dbresolver.Register(dbresolver.Config{
		// Sources: []gorm.Dialector{mysql.Open(address)},
		Replicas: []gorm.Dialector{mysql.Open(address)},
	}))
	if err != nil {
		logger.Panic("Connect database faild")
	}
}

func GetDB() *gorm.DB {
	return db
}
