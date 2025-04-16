package main

import (
	"log"
	"order-service/internal/adapter/http"
	"order-service/internal/infrastructure/db"
	"order-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := db.NewOrderRepo("database.db")
	uc := usecase.NewOrderUsecase(repo)
	h := http.NewOrderHandler(uc)

	r := gin.Default()
	r.POST("/orders", h.CreateOrder)
	r.GET("/orders/:id", h.GetOrderByID)
	log.Fatal(r.Run(":8082"))
}
