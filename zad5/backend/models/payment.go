package models

type Payment struct {
    ID        int    `json:"id"`
    CartID    int    `json:"cartId"`
    Status    string `json:"status"`
    Timestamp string `json:"timestamp"`
}
