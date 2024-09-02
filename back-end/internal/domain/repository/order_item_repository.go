package repository

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/entity/dto"
)

type OrderItemRepository interface {
    FindById(id string) (*entity.OrderItem, error)
    Create(store *entity.OrderItem) error
    Update(store *entity.OrderItem) (*entity.OrderItem, error)
    Delete(id string) error
    FindAllOrderItemsByOrderId(storeId string) ([]*dto.OrderItemWithProductDTO ,error) 
}