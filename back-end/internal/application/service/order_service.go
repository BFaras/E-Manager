package service

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/db"
	"database/sql"
)

type OrderService struct {
    repository repository.OrderRepository
}

func NewOrderService(database *sql.DB) *OrderService {
    return &OrderService{
        repository: db.NewOrderRepository(database),
    }
}

func (s *OrderService) GetOrder(id string) (*entity.Order, error) {
    order, err := s.repository.FindByID(id)
    if err != nil {
        return nil, err
    }
    return order, nil
}