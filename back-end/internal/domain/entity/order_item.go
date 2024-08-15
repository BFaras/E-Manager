package entity

type OrderItem struct {
    Id        string  `json:"id"`
    OrderId   string  `json:"orderId"`
    ProductId string  `json:"productId"`
}