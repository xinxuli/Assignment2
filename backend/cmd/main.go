package main

import (
	"log"
	"backend/handler"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	handler.InitGin(router)
	log.Println("Server is running on port http://127.0.0.1:8080")
	router.Run("127.0.0.1:8080")
}