package entity

import "time"

type Order struct {
    ID         string      `json:"id"`
    StoreID    string      `json:"storeId"`
    Store      Store       `json:"store"`
    OrderItems []OrderItem `json:"orderItems"`
    IsPaid     bool        `json:"isPaid"`
    Phone      string      `json:"phone"`
    Address    string      `json:"address"`
    CreatedAt  time.Time   `json:"createdAt"`
    UpdatedAt  time.Time   `json:"updatedAt"`
}