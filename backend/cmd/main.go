package main

import (
	"backend/handler"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	handler.InitGin(router)
	log.Println("Server is running on port http://127.0.0.1:8080")
	router.Run("127.0.0.1:8080")
}
