package handlers

import (
    "net/http"
    "sync"
    "strconv"
    "backend/models"
    "github.com/gin-gonic/gin"
)

var (
    carts  = map[int]models.Cart{}
    cartID = 1
    muCart = sync.Mutex{}
)

func CreateCart(c *gin.Context) {
    var cart models.Cart
    if err := c.ShouldBindJSON(&cart); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart data"})
        return
    }

    total := 0.0
    for _, item := range cart.Items {
        total += item.Product.Price * float64(item.Quantity)
    }
    cart.Total = total

    muCart.Lock()
    cart.ID = cartID
    carts[cartID] = cart
    cartID++
    muCart.Unlock()

    c.JSON(http.StatusCreated, cart)
}

func GetCart(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
        return
    }

    muCart.Lock()
    cart, exists := carts[id]
    muCart.Unlock()

    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
        return
    }

    c.JSON(http.StatusOK, cart)
}

func AddProductToCart(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
        return
    }

    var input struct {
        Product  models.Product `json:"product"`
        Quantity int            `json:"quantity"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    muCart.Lock()
    cart, exists := carts[id]
    if !exists {
        muCart.Unlock()
        c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
        return
    }

    // check if product already in cart
    found := false
    for i, item := range cart.Items {
        if item.Product.ID == input.Product.ID {
            cart.Items[i].Quantity += input.Quantity
            found = true
            break
        }
    }

    if !found {
        cart.Items = append(cart.Items, models.CartItem{
            Product:  input.Product,
            Quantity: input.Quantity,
        })
    }

    // recalculate total
    total := 0.0
    for _, item := range cart.Items {
        total += item.Product.Price * float64(item.Quantity)
    }
    cart.Total = total

    carts[id] = cart
    muCart.Unlock()

    c.JSON(http.StatusOK, cart)
}
