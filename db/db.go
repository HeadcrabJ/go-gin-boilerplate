package db

import (
	"github.com/dajeo/go-gin-boilerplate/config"
	"github.com/rs/zerolog/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	c := config.Get()
	dsn := "host=" + c.GetString("pg.host") +
		" user=" + c.GetString("pg.user") +
		" password=" + c.GetString("pg.pass") +
		" dbname=" + c.GetString("pg.name") +
		" sslmode=require"

	conn, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	db = conn
}

func Get() *gorm.DB {
	return db
}
