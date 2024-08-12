package entity

import "time"

type Color struct {
    ID        string    `json:"id"`
    StoreID   string    `json:"storeId"`
    Store     Store     `json:"store"`
    Products  []Product `json:"products"`
    Name      string    `json:"name"`
    Value     string    `json:"value"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}