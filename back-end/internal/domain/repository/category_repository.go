package repository

import ("back-end/internal/domain/entity")

type CategoryRepository interface {
    FindByID(id string) (*entity.Category, error)
    Create(store *entity.Category) error
    Update(store *entity.Category) (*entity.Category, error)
    Delete(id string) error
}
