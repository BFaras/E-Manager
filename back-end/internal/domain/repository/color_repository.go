package repository

import "back-end/internal/domain/entity"

type ColorRepository interface {
    FindById(id string) (*entity.Color, error)
    Create(store *entity.Color) error
    Update(store *entity.Color) (error)
    Delete(id string) error
    FindAllColors() ([]*entity.Color, error)
}