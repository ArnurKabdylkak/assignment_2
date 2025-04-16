package main

import (
	"log"

	"api-gateway/internal/adapter/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router := http.NewRouter()
	r.Any("/*proxyPath", func(c *gin.Context) {
		router.HandleContext(c)
	})
	log.Fatal(r.Run(":8080"))
}
