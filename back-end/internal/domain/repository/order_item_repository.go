package repository

import "back-end/internal/domain/entity"

type OrderItemRepository interface {
    FindByID(id string) (*entity.OrderItem, error)
    Create(store *entity.OrderItem) error
    Update(store *entity.OrderItem) (*entity.OrderItem, error)
    Delete(id string) error
}