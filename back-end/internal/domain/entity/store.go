package entity

import ("time")

type Store struct {
	Id string `json:"id"`
	Name string `json:"name"`
	UserId string `json:"userId"`
	Billboards []Billboard `json:"billboards"`
	Categories []Category `json:"categories"`
	Sizes []Size `json:"sizes"`
	Colors []Color `json:"colors"`
	Products []Product `json:"products"`
	Orders []Order `json:"orders"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}