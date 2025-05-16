package handlers

import (
	"net/http"
	"backend/models"
	"github.com/gin-gonic/gin"
)

func ProcessPayment(c *gin.Context) {
	var payment models.PaymentRequest

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	cart, exists := carts[payment.CartID]
	if !exists || len(cart.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found or empty"})
		return
	}

	if len(payment.CardNumber) < 12 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card number"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Payment processed successfully",
		"cartId":   payment.CartID,
		"total":    calculateCartTotal(cart),
		"last4":    payment.CardNumber[len(payment.CardNumber)-4:],
	})
}

func calculateCartTotal(cart models.Cart) float64 {
	total := 0.0
	for _, item := range cart.Items {
		for _, product := range products {
			if product.ID == item.ProductID {
				total += float64(item.Quantity) * product.Price
			}
		}
	}
	return total
}
