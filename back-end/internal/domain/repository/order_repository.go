package repository

import (
	"back-end/internal/domain/entity"
)

type OrderRepository interface {
    FindByID(id string) (*entity.Order, error)
    Create(store *entity.Order) error
    Update(store *entity.Order) (error)
    Delete(id string) error
}