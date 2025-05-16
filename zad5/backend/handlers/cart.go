package handlers

import (
	"net/http"
	"strconv"
	"backend/models"
	"github.com/gin-gonic/gin"
)

var (
	carts  = make(map[int]models.Cart)
	nextID = 1
)

func CreateCart(c *gin.Context) {
	cart := models.Cart{
		ID:    nextID,
		Items: []models.CartItem{},
	}

	carts[nextID] = cart
	nextID++

	c.JSON(http.StatusCreated, cart)
}

func GetCart(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	cart, exists := carts[id]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}
	c.JSON(http.StatusOK, cart)
}

func AddToCart(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input models.CartItem

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	cart, exists := carts[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	found := false
	for i := range cart.Items {
		if cart.Items[i].ProductID == input.ProductID {
			cart.Items[i].Quantity += input.Quantity
			found = true
			break
		}
	}
	if !found {
		cart.Items = append(cart.Items, input)
	}

	carts[id] = cart
	c.JSON(http.StatusOK, cart)
}

func ClearCart(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, exists := carts[id]
	if exists {
		carts[id] = models.Cart{ID: id, Items: []models.CartItem{}}
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart cleared"})
}

func RemoveItemFromCart(c *gin.Context) {
	cartID, _ := strconv.Atoi(c.Param("id"))
	productID, _ := strconv.Atoi(c.Param("productId"))

	cart, exists := carts[cartID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	newItems := []models.CartItem{}
	removed := false

	for _, item := range cart.Items {
		if item.ProductID == productID {
			removed = true
			continue 
		}
		newItems = append(newItems, item)
	}

	if !removed {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found in cart"})
		return
	}

	cart.Items = newItems
	carts[cartID] = cart

	c.JSON(http.StatusOK, cart)
}
