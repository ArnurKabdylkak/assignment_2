package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("supersecretkey")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {
	var creds Credentials

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if creds.Username != "admin" || creds.Password != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": creds.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.POST("/login", LoginHandler)

	protected := r.Group("/")
	protected.Use(AuthMiddleware())
	{
		protected.POST("/products", proxy("http://inventory-service:8081"))
		protected.GET("/products/:id", proxy("http://inventory-service:8081"))
		protected.POST("/orders", proxy("http://order-service:8082"))
		protected.GET("/orders/:id", proxy("http://order-service:8082"))
	}

	return r
}

func proxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		client := http.DefaultClient
		req, _ := http.NewRequest(c.Request.Method, target+c.Request.RequestURI, c.Request.Body)
		req.Header = c.Request.Header
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()
		c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
	}
}
