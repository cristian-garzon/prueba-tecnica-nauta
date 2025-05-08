package routes

import (
	"prueba-tecnica-nauta/app/infrastructure/builder"
	"prueba-tecnica-nauta/app/infrastructure/server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupWriteRoutes(router *gin.RouterGroup, actions *builder.Actions) {
	writeHandler := handlers.NewWriteHandler(actions)
	router.POST("/bookings", writeHandler.CreateBooking)
}
