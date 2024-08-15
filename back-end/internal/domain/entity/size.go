package entity

import ("time")

type Size struct {
	Id string `json:"id"`
	StoreId string `json:"storeId"`
	Name string `json:"name"`
	Value string `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

