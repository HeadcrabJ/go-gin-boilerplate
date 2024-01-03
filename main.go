package main

import (
	"flag"
	"github.com/dajeo/go-gin-boilerplate/config"
	"github.com/dajeo/go-gin-boilerplate/db"
	"github.com/dajeo/go-gin-boilerplate/middlewares"
	"github.com/dajeo/go-gin-boilerplate/rdb"
	"github.com/dajeo/go-gin-boilerplate/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

func main() {
	initLogger()

	// Generate docs (swag init) and uncomment line below
	// docs.SwaggerInfo.BasePath = "/"

	cfg := *flag.String("cfg", "dev", "config file to use")
	flag.Parse()

	log.Info().Msg("Initializing configuration")
	config.Init(cfg)

	log.Info().Msg("Initializing database")
	db.Init()

	log.Info().Msg("Initializing redis")
	rdb.Init()

	log.Info().Msg("Initializing middlewares")
	middlewares.InitJWT()

	log.Info().Msg("Starting server")
	server.Run()
}

func initLogger() {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "02 Jan 15:04:05",
	}).With().Caller().Logger()
}
