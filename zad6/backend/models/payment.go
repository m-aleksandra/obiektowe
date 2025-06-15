package models

type PaymentRequest struct {
	CartID     int    `json:"cartId"`
	CardNumber string `json:"cardNumber"`
}
