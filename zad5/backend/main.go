// main.go
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
        AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        AllowCredentials: true,
    }))

    r.GET("/api/products", handlers.GetProducts)
    r.POST("/api/products", handlers.AddProduct)

    r.POST("/api/cart", handlers.CreateCart)
    r.GET("/api/cart/:id", handlers.GetCart)
	r.PATCH("/api/cart/:id", handlers.AddProductToCart)

    r.POST("/api/payment", handlers.SubmitPayment)

    r.Run(":8080")
}