package main

import (
	"api-gateway/internal/adapter/http"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router := http.NewRouter()
	r.Any("/*proxyPath", router) // proxy all to internal router
	log.Fatal(r.Run(":8080"))
}
