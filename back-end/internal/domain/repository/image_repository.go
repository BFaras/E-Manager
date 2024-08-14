package repository

import ("back-end/internal/domain/entity")

type ImageRepository interface {
    FindByID(id string) (*entity.Image, error)
    Create(store *entity.Image) error
    Update(store *entity.Image) (*entity.Image, error)
    Delete(id string) error

}
