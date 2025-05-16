package models

type CartItem struct {
	ProductID int `json:"productId"`
	Quantity  int `json:"quantity"`
}

type Cart struct {
	ID    int        `json:"id"`
	Items []CartItem `json:"items"`
}
