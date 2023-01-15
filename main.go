package main

import (
	"github.com/HeadcrabJ/go-gin-boilerplate/config"
	"github.com/HeadcrabJ/go-gin-boilerplate/db"
	"github.com/HeadcrabJ/go-gin-boilerplate/middlewares"
	"github.com/HeadcrabJ/go-gin-boilerplate/server"
)

func main() {
	// Generate docs (swag init) and uncomment line below
	// docs.SwaggerInfo.BasePath = "/"

	config.Init("dev")

	db.InitDB()
	db.InitRedis()

	middlewares.InitJWT()

	server.Init()
}
