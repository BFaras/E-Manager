package repository

import ("back-end/internal/domain/entity")

type ImageRepository interface {
    FindById(id string) (*entity.Image, error)
    Create(store *entity.Image) error
    Update(store *entity.Image) (*entity.Image, error)
    Delete(id string) error

}
