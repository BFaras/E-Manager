package repository

import "back-end/internal/domain/entity"

type SizeRepository interface {
    FindByID(id string) (*entity.Size, error)
    Create(size *entity.Size) (error)
    Update(size *entity.Size) (error)
    Delete(id string) error
    FindAllSizes() ([]*entity.Size, error)
}