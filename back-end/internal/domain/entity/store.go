package entity

import ("time")

type Store struct {
	Id string `json:"id"`
	Name string `json:"name"`
	UserId string `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}