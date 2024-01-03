package rdb

import (
	"github.com/dajeo/go-gin-boilerplate/config"
	"github.com/go-redis/redis/v9"
)

var rdb *redis.Client

func Init() {
	c := config.Get()
	rdb = redis.NewClient(&redis.Options{
		Addr:     c.GetString("redis.host"),
		Password: c.GetString("redis.pass"),
		DB:       0,
	})
}

func Get() *redis.Client {
	return rdb
}
