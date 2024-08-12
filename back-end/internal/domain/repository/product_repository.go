package repository

import ("back-end/internal/domain/entity")

type ProductRepository interface {
    FindByID(id string) (*entity.Product, error)
    Save(product *entity.Product) error
}
