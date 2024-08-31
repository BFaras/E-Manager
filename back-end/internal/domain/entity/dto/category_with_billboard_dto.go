package dto

import ("back-end/internal/domain/entity"
"time")

type CategoryWithBillboardDTO struct {
	Id          string    `json:"id"`
	StoreId     string    `json:"storeId"`
	Billboard   *entity.Billboard    `json:"billboard"`
	BillboardId string    `json:"billboardId"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}