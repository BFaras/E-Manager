package entity

import ("time")

type Store struct {
	ID string `json:"id"`
	Name string `json:"name"`
	UserID string `json:"userId"`
	Billboards []Billboard `json:"billboards"`
	Categories []Category `json:"categories"`
	Sizes []Size `json:"sizes"`
	Colors []Color `json:"colors"`
	Products []Product `json:"products"`
	Orders []Order `json:"orders"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}