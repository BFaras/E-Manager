package entity

import "time"

type Category struct {
    Id        string    `json:"id"`
    StoreId   string    `json:"storeId"`
    BillboardId string  `json:"billboardId"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}