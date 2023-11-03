package main

import (
	"github.com/gin-gonic/gin"
	"go-rms/database"
	"go-rms/middleware"
	"go-rms/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	
	router.Use(middleware.Authentication())

	routes.UserRoutes(router)
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoicemRoutes(router)

	router.Run(":" + port)

}
