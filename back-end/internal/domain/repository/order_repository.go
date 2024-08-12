package repository

import "back-end/internal/domain/entity"

type OrderRepository interface {
    FindByID(id string) (*entity.Order, error)
    Save(order *entity.Order) error
    // Add other methods as needed
}