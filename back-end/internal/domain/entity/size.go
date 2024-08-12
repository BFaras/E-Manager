package entity

type Size struct {
	ID string `json:"id"`
	StoreID string `json:"store_id"`
	Store Store `json:"store"`
	Products []Product `json:"products"`
	Name string `json:"name"`
	Value string `json:"value"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

