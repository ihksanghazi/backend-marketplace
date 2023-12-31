package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/database"
	"github.com/ihksanghazi/backend-marketplace/model/domain"
	"github.com/ihksanghazi/backend-marketplace/routers"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	// middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()
	database.DB.AutoMigrate(domain.User{}, domain.Product{}, domain.Review{})
	// database.Seeder()

	routers.TestRouter(r.Group("/api/test"))
	routers.RegionRouter(r.Group("/api/region"))
	routers.UserRouter(r.Group("/api/user"))
	routers.StoreRouter(r.Group("/api/store"))
	routers.ProductRouter(r.Group("/api/product"))
	routers.CartRouter(r.Group("/api/cart"))
	routers.TransactionRouter(r.Group("/api/transaction"))
	routers.ReviewRouter(r.Group("/api/review"))

	r.Run(":5000")
}
