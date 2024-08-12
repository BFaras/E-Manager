package repository

import ("back-end/internal/domain/entity")

type ImageRepository interface {
    FindByID(id string) (*entity.Image, error)
    Save(image *entity.Image) error

}
