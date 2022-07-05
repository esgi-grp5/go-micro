package main

import (
	"go-micro/internal/config"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

// server struct
type server struct {
	gin    *gin.Engine
	config config.Configuration
	oauth  config.OAuthApp
}

// @title Gin Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func newServer(c config.Configuration) *server {
	// Initialize server
	s := &server{
		gin:    gin.New(),
		config: c,
		oauth:  c.OAuthApp,
	}
	// Initialize router
	s.routes()
	return s
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UTC().UnixNano())
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}