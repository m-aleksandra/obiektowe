package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "backend/handlers"
)

func main() {
    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PATCH", "PUT"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        AllowCredentials: true,
    }))

    r.GET("/api/products", handlers.GetProducts)
    r.GET("/api/products/:id", handlers.GetProductByID)
	r.POST("/api/cart", handlers.CreateCart)
	r.GET("/api/cart/:id", handlers.GetCart)
	r.POST("/api/cart/:id/add", handlers.AddToCart)
	r.DELETE("/api/cart/:id", handlers.ClearCart)
	r.DELETE("/api/cart/:id/item/:productId", handlers.RemoveItemFromCart)
	r.POST("/api/payment", handlers.ProcessPayment)

    r.Run(":8080")
}