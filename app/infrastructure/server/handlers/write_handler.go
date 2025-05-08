package handlers

import (
	"net/http"
	"prueba-tecnica-nauta/app/domain/model"
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"prueba-tecnica-nauta/app/infrastructure/builder"

	"github.com/gin-gonic/gin"
)

type WriteHandler struct {
	actions *builder.Actions
}

func NewWriteHandler(actions *builder.Actions) *WriteHandler {
	return &WriteHandler{actions: actions}
}

func (h *WriteHandler) CreateBooking(c *gin.Context) {

	var booking Dto.BookingDto
	if err := c.ShouldBindJSON(&booking); err != nil {
		model.HandleError(c.Writer, model.NewAppError("invalid body "+err.Error(), model.InvalidBodyError))
		return
	}

	err := h.actions.AddBookingAction.Execute(booking)
	if err != nil {
		model.HandleError(c.Writer, err)
		return
	}
	c.JSON(http.StatusOK, booking)
}
