package entity

import ("time")

type Billboard struct {
	Id string `json:"id"`
	StoreId string `json:"storeId"`
	Label string `json:"label"`
	ImageUrl string `json:"imageUrl"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsActive bool `json:"isActive"`
}