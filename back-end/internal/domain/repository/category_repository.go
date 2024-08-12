package repository

import ("back-end/internal/domain/entity")

type CategoryRepository interface {
    FindByID(id string) (*entity.Category, error)
    Save(category *entity.Category) error
    // Add other methods as needed
}
