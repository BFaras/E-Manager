package service

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/entity/dto"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/db"
	"database/sql"
)

type OrderService struct {
    orderRepository      repository.OrderRepository
    orderItemRepository  repository.OrderItemRepository
    productRepository    repository.ProductRepository
}

func NewOrderService(database *sql.DB) *OrderService {
    return &OrderService{
        orderRepository: db.NewOrderRepository(database),
        orderItemRepository: db.NewOrderItemRepository(database),
        productRepository: db.NewProductRepository(database),
    }
}

func (s *OrderService) GetOrder(id string) (*entity.Order, error) {
    order, err := s.orderRepository.FindById(id)
    if err != nil {
        return nil, err
    }
    return order, nil
}

func (s *OrderService) GetAllOrdersWithExtraInformationByStoreId(storeId string) ([]*dto.OrderWithExtraInfoDTO, error) {

    orders, err := s.orderRepository.FindAllOrdersWithExtraInfoByStoreId(storeId)
    if err != nil {
        return nil, err
    }

    if orders == nil{
        return nil, nil
    }

    for _,order := range orders {
        orderItems, err := s.orderItemRepository.FindAllOrderItemsByOrderId(order.Id)
        if err != nil {
            return nil, err
        }
        order.OrderItems = orderItems
        
        if (order.OrderItems != nil) {
            for _, orderItem := range order.OrderItems {
                product, err := s.productRepository.FindById(orderItem.ProductId)
                if err!= nil {
                    return nil, err
                }
                orderItem.Product = product
            }
        }

    }

    return orders, nil

}