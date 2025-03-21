package dto

type ProductFilterDTO struct {
    ColorId    string `json:"colorId"`
    SizeId     string `json:"sizeId"`
    CategoryId string `json:"categoryId"`
    IsFeatured string `json:"isFeatured"`
}
