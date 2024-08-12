package entity

import "time"

type Image struct {
    ID        string    `json:"id"`
    ProductID string    `json:"productId"`
    Product   Product   `json:"product"`
    URL       string    `json:"url"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}