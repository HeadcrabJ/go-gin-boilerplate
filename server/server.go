package server

import (
	"github.com/HeadcrabJ/go-gin-boilerplate/config"
	"github.com/gin-contrib/cors"
)

func Init() {
	r := NewRouter()

	// See this before deploy to production
	// https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies
	proxyErr := r.SetTrustedProxies(nil)
	if proxyErr != nil {
		return
	}

	r.Use(cors.Default()) // Configure before production

	err := r.Run(":" + config.GetConfig().GetString("server.port"))
	if err != nil {
		return
	}
}
