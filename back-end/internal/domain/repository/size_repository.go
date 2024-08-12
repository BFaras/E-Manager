package repository

import "back-end/internal/domain/entity"

type SizeRepository interface {
    FindByID(id string) (*entity.Size, error)
    Save(size *entity.Size) error
    // Add other methods as needed
}