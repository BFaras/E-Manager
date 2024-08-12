package repository

import "back-end/internal/domain/entity"

type ColorRepository interface {
    FindByID(id string) (*entity.Color, error)
    Save(color *entity.Color) error
}