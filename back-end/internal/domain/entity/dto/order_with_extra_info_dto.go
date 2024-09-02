package dto

import (
	"time"
)

type OrderWithExtraInfoDTO struct {
    Id         string      `json:"id"`
    StoreId    string      `json:"storeId"`
    IsPaid     bool        `json:"isPaid"`
    Phone      string      `json:"phone"`
    Address    string      `json:"address"`
	OrderItems []*OrderItemWithProductDTO `json:"orderItems"`
    CreatedAt  time.Time   `json:"createdAt"`
    UpdatedAt  time.Time   `json:"updatedAt"`
}