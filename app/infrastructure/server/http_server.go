package server

import (
	"net/http"
	"prueba-tecnica-nauta/app/infrastructure/builder"
	"prueba-tecnica-nauta/app/infrastructure/config"
	"prueba-tecnica-nauta/app/infrastructure/server/routes"

	"github.com/gin-gonic/gin"
)

func SetupServer(config *config.ServerConfig, actions *builder.Actions) *http.Server {

	ginServer := gin.New()

	prefix := ginServer.Group(config.Prefix)

	prefix.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	routes.SetupReadRoutes(prefix, actions)
	routes.SetupWriteRoutes(prefix, actions)

	return &http.Server{
		Addr:         ":" + config.Port,
		Handler:      ginServer,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}
}
