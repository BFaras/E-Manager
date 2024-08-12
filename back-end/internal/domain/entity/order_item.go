package entity

type OrderItem struct {
    ID        string  `json:"id"`
    OrderID   string  `json:"orderId"`
    Order     Order   `json:"order"`
    ProductID string  `json:"productId"`
    Product   Product `json:"product"`
}