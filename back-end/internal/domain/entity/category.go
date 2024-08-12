package entity

import "time"

type Category struct {
    ID        string    `json:"id"`
    StoreID   string    `json:"storeId"`
    Store     Store     `json:"store"`
    BillboardID string  `json:"billboardId"`
    Billboard Billboard `json:"billboard"`
    Products  []Product `json:"products"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}