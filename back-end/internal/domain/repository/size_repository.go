package repository

import "back-end/internal/domain/entity"

type SizeRepository interface {
    FindByID(id string) (*entity.Size, error)
    Create(store *entity.Size) error
    Update(store *entity.Size) (*entity.Size, error)
    Delete(id string) error
}