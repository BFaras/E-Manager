package repository

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/entity/dto"
)

type ProductRepository interface {
    FindById(id string) (*entity.Product, error)
    Create(product *dto.ProductWithImageDTO) error
    Update(product *entity.Product) (error)
    Delete(id string) error
    FindAllProductsWithExtraInformationByStoreId(storeId string) ([]*dto.ProductWithExtraInfoDTO, error)
    FindAllProductsWithImageById(id string) (*dto.ProductWithImageDTO, error)
}
