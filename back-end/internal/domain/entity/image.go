package entity

import "time"

type Image struct {
    Id       string    `json:"id"`
    ProductId string    `json:"productId"`
    URL       string    `json:"url"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}