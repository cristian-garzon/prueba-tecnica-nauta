package handlers

import (
	"net/http"
	"prueba-tecnica-nauta/app/domain/model"
	"prueba-tecnica-nauta/app/infrastructure/builder"

	"github.com/gin-gonic/gin"
)

type ReadHandler struct {
	actions *builder.Actions
}

func NewReadHandler(actions *builder.Actions) *ReadHandler {
	return &ReadHandler{actions: actions}
}

func (h *ReadHandler) GetContainersByBookingId(c *gin.Context) {
	id := c.Param("id")
	containers, err := h.actions.GetContainersByBookingIdAction.Execute(id)
	if err != nil {
		model.HandleError(c.Writer, err)
		return
	}
	c.JSON(http.StatusOK, containers)
}

func (h *ReadHandler) GetContainersByOrderId(c *gin.Context) {
	id := c.Param("id")
	containers, err := h.actions.GetContainersByOrderIdAction.Execute(id)
	if err != nil {
		model.HandleError(c.Writer, err)
		return
	}
	c.JSON(http.StatusOK, containers)
}

func (h *ReadHandler) GetContainersByEmail(c *gin.Context) {
	email := c.Param("email")
	containers, err := h.actions.GetContainersByEmailAction.Execute(email)
	if err != nil {
		model.HandleError(c.Writer, err)
		return
	}
	c.JSON(http.StatusOK, containers)
}

func (h *ReadHandler) GetOrdersByBookingId(c *gin.Context) {
	id := c.Param("id")
	orders, err := h.actions.GetOrdersByBookingIdAction.Execute(id)
	if err != nil {
		model.HandleError(c.Writer, err)
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *ReadHandler) GetOrdersByContainerId(c *gin.Context) {
	id := c.Param("id")
	orders, err := h.actions.GetOrdersByContainerIdAction.Execute(id)
	if err != nil {
		model.HandleError(c.Writer, err)
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *ReadHandler) GetOrdersByEmail(c *gin.Context) {
	email := c.Param("email")
	orders, err := h.actions.GetOrdersByEmailAction.Execute(email)
	if err != nil {
		model.HandleError(c.Writer, err)
		return
	}
	c.JSON(http.StatusOK, orders)
}
