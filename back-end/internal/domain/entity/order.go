package entity

import "time"

type Order struct {
    Id         string      `json:"id"`
    StoreId    string      `json:"storeId"`
    IsPaid     bool        `json:"isPaid"`
    Phone      string      `json:"phone"`
    Address    string      `json:"address"`
    CreatedAt  time.Time   `json:"createdAt"`
    UpdatedAt  time.Time   `json:"updatedAt"`
}