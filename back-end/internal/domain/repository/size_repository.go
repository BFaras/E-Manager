package repository

import "back-end/internal/domain/entity"

type SizeRepository interface {
    FindById(id string) (*entity.Size, error)
    Create(size *entity.Size) (error)
    Update(size *entity.Size) (error)
    Delete(id string) error
    FindAllSizes(storeId string) ([]*entity.Size, error)
}