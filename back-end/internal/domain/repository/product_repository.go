package repository

import ("back-end/internal/domain/entity")

type ProductRepository interface {
    FindByID(id string) (*entity.Product, error)
    Create(store *entity.Product) error
    Update(store *entity.Product) (error)
    Delete(id string) error
}
