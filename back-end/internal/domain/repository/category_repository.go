package repository

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/entity/dto"
)

type CategoryRepository interface {
    FindByID(id string) (*entity.Category, error)
    Create(store *entity.Category) error
    Update(store *entity.Category) (error)
    Delete(id string) error
    FindCategoriesWithBillboard(storeId string) ([]*dto.CategoryWithBillboardDTO, error)
}
