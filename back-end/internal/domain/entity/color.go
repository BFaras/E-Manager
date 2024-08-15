package entity

import "time"

type Color struct {
    Id        string    `json:"id"`
    StoreId   string    `json:"storeId"`
    Name      string    `json:"name"`
    Value     string    `json:"value"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}