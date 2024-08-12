package entity

import "time"

type Product struct {
    ID         string    `json:"id"`
    StoreID    string    `json:"storeId"`
    Store      Store     `json:"store"`
    CategoryID string    `json:"categoryId"`
    Category   Category  `json:"category"`
    Name       string    `json:"name"`
    Price      float64   `json:"price"`
    IsFeatured bool      `json:"isFeatured"`
    IsArchived bool      `json:"isArchived"`
    SizeID     string    `json:"sizeId"`
    Size       Size      `json:"size"`
    ColorID    string    `json:"colorId"`
    Color      Color     `json:"color"`
    Images     []Image   `json:"images"`
    OrderItems []OrderItem `json:"orderItems"`
    CreatedAt  time.Time `json:"createdAt"`
    UpdatedAt  time.Time `json:"updatedAt"`
}