package server

import (
	"github.com/dajeo/go-gin-boilerplate/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
)

func Run() {
	gin.DefaultWriter = io.Discard

	r := NewRouter()

	// See this before deploy to production
	// https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies
	proxyErr := r.SetTrustedProxies(nil)
	if proxyErr != nil {
		return
	}

	r.Use(cors.Default()) // Configure before production

	err := r.Run(":" + config.Get().GetString("server.port"))
	if err != nil {
		return
	}
}
