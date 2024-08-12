package repository

import "back-end/internal/domain/entity"

type OrderItemRepository interface {
    FindByID(id string) (*entity.OrderItem, error)
    Save(orderItem *entity.OrderItem) error
}