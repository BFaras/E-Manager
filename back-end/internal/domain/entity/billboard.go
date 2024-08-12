package entity

import ("time")

type Billboard struct {
	ID string `json:"id"`
	StoreID string `json:"storeId"`
	Store Store `json:"store"`
	Label string `json:"label"`
	ImageURL string `json:"imageUrl"`
	Categories []Category `json:"categories"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}