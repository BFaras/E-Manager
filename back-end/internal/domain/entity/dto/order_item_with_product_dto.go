package dto

import "back-end/internal/domain/entity"

type OrderItemWithProductDTO struct {
    Id        string  `json:"id"`
    OrderId   string  `json:"orderId"`
    ProductId string  `json:"productId"`
	Product *entity.Product `json:"product"`
}