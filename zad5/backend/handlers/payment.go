package handlers

import (
    "net/http"
    "time"
    "sync"
    "backend/models"
    "github.com/gin-gonic/gin"
)

var (
    payments   = []models.Payment{}
    paymentID  = 1
    muPayment  = sync.Mutex{}
)

func SubmitPayment(c *gin.Context) {
    var req struct {
        CartID int `json:"cartId"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment request"})
        return
    }

    muCart.Lock()
    cart, exists := carts[req.CartID]
    muCart.Unlock()

    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
        return
    }

    muPayment.Lock()
    payment := models.Payment{
        ID:        paymentID,
        CartID:    cart.ID,
        Status:    "success",
        Timestamp: time.Now().Format(time.RFC3339),
    }
    payments = append(payments, payment)
    paymentID++
    muPayment.Unlock()

    c.JSON(http.StatusCreated, payment)
}
