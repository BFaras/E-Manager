package repository

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/entity/dto"
)

type OrderRepository interface {
    FindById(id string) (*entity.Order, error)
    Create(store *entity.Order) error
    Update(store *entity.Order) (error)
    Delete(id string) error
    FindAllOrdersWithExtraInfoByStoreId(storeId string) ([]*dto.OrderWithExtraInfoDTO ,error) 
}