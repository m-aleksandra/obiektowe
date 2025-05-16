package handlers

import (
	"backend/models"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

var products = []models.Product{
    {ID: 1, Name: "Coffee", Price: 9.99},
    {ID: 2, Name: "Tea", Price: 4.99},
    {ID: 3, Name: "Juice", Price: 5.49},
}

func GetProducts(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, product := range products {
		if product.ID == id {
			c.IndentedJSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

