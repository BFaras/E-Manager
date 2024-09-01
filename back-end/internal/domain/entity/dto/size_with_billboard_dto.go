package dto

import ("back-end/internal/domain/entity"
"time")

type SizeWithBillboardDTO struct {
	Id          string     `json:"id"`
	StoreId     string     `json:"storeId"`
	Store   *entity.Store  `json:"store"`
	Name        string     `json:"name"`
	Value		string     `json:"value"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}
