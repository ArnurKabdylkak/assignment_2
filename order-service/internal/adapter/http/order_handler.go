package http

import (
	"net/http"
	"order-service/internal/domain"
	"order-service/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct{ uc usecase.OrderUsecase }

func NewOrderHandler(u usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{uc: u}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var o domain.Order
	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.CreateOrder(&o); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, o)
}

func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	o, err := h.uc.GetOrderByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, o)
}
