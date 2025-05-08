package routes

import (
	"prueba-tecnica-nauta/app/infrastructure/builder"
	"prueba-tecnica-nauta/app/infrastructure/server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupReadRoutes(router *gin.RouterGroup, actions *builder.Actions) {
	readHandler := handlers.NewReadHandler(actions)

	router.GET("/containers/booking/:id", readHandler.GetContainersByBookingId)
	router.GET("/containers/order/:id", readHandler.GetContainersByOrderId)
	router.GET("/containers/email/:email", readHandler.GetContainersByEmail)
	router.GET("/orders/booking/:id", readHandler.GetOrdersByBookingId)
	router.GET("/orders/container/:id", readHandler.GetOrdersByContainerId)
	router.GET("/orders/email/:email", readHandler.GetOrdersByEmail)
}
