package repository

import "back-end/internal/domain/entity"

type ColorRepository interface {
    FindByID(id string) (*entity.Color, error)
    Create(store *entity.Color) error
    Update(store *entity.Color) (*entity.Color, error)
    Delete(id string) error

}