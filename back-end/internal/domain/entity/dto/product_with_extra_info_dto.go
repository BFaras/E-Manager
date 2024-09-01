package dto


import ("back-end/internal/domain/entity"
"time")

type ProductWithExtraInfoDTO struct {
	Id          string    `json:"id"`
    Name       string  `json:"name"`
    Price      float64 `json:"price"`
    IsFeatured bool    `json:"isFeatured"`
    IsArchived bool    `json:"isArchived"`
    StoreId    string  `json:"storeId"`
    CategoryId string  `json:"categoryId"`
    SizeId     string  `json:"sizeId"`
    ColorId    string  `json:"colorId"`
    Category   *entity.Category    `json:"category"`
    Color      *entity.Color    `json:"color"`
    Size       *entity.Size    `json:"size"`
    CreatedAt  time.Time `json:"createdAt"`
    UpdatedAt  time.Time `json:"updatedAt"`
}