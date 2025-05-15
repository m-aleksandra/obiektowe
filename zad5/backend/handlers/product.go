package handlers

import (
    "net/http"
    "sync"
    "backend/models"
    "github.com/gin-gonic/gin"
)

var (
    products = []models.Product{
        {ID: 1, Name: "Kawa", Price: 12.99},
        {ID: 2, Name: "Herbata", Price: 8.50},
        {ID: 3, Name: "Sok", Price: 6.00},
    }
    
    productID = 4
    muProduct = sync.Mutex{}
)

func GetProducts(c *gin.Context) {
    c.JSON(http.StatusOK, products)
}

func AddProduct(c *gin.Context) {
    var p models.Product
    if err := c.ShouldBindJSON(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product"})
        return
    }

    muProduct.Lock()
    p.ID = productID
    productID++
    products = append(products, p)
    muProduct.Unlock()

    c.JSON(http.StatusCreated, p)
}
