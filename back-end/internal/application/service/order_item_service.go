package service

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/db"
	"database/sql"
)

type OrderItemService struct {
    repository repository.OrderItemRepository
}

func NewOrderItemService(database *sql.DB) *OrderItemService {
    return &OrderItemService{
        repository: db.NewOrderItemRepository(database),
    }
}

func (s *OrderItemService) GetOrderItem(id string) (*entity.OrderItem, error) {
    orderItem, err := s.repository.FindById(id)
    if err != nil {
        return nil, err
    }
    return orderItem, nil
}