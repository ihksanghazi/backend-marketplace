package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/routers"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	routers.UserRouter(r.Group("/api/user"))

	r.Run(":5000")
}
