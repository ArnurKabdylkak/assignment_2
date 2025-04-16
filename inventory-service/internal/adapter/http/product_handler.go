package http

import (
	"inventory-service/internal/domain"
	"inventory-service/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct{ uc usecase.ProductUsecase }

func NewProductHandler(u usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{uc: u}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var p domain.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.CreateProduct(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, p)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p, err := h.uc.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}
