package entity

import "time"

type Product struct {
    Id         string    `json:"id"`
    StoreId    string    `json:"storeId"`
    CategoryId string    `json:"categoryId"`
    Name       string    `json:"name"`
    Price      float64   `json:"price"`
    IsFeatured bool      `json:"isFeatured"`
    IsArchived bool      `json:"isArchived"`
    SizeID     string    `json:"sizeId"`
    ColorID    string    `json:"colorId"`
    CreatedAt  time.Time `json:"createdAt"`
    UpdatedAt  time.Time `json:"updatedAt"`
}