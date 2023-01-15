package db

import (
	"github.com/HeadcrabJ/go-gin-boilerplate/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	c := config.GetConfig()
	dsn := "host=" + c.GetString("pg.host") +
		" user=" + c.GetString("pg.user") +
		" password=" + c.GetString("pg.pass") +
		" dbname=" + c.GetString("pg.name") +
		" sslmode=require"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
}

func GetDB() *gorm.DB {
	return db
}
