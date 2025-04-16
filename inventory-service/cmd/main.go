package main

import (
	"inventory-service/internal/adapter/http"
	repo "inventory-service/internal/infrastructure/db"
	"inventory-service/internal/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repo.NewProductRepo("database.db")
	uc := usecase.NewProductUsecase(repo)
	h := http.NewProductHandler(uc)

	r := gin.Default()
	r.POST("/products", h.CreateProduct)
	r.GET("/products/:id", h.GetProductByID)
	log.Fatal(r.Run(":8081"))
}
