package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/entity/dto"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/logger"
	"database/sql"

	"go.uber.org/zap"
)

type orderRepository struct {
    db *sql.DB
}

func NewOrderRepository(db *sql.DB) repository.OrderRepository {
    return &orderRepository{db: db}
}

func (r *orderRepository) FindById(id string) (*entity.Order ,error) {
    order:= &entity.Order{}
    query := `SELECT * FROM "public"."Order" WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(&order.Id, &order.StoreId,&order.IsPaid ,&order.Phone,&order.Address, &order.CreatedAt, &order.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return order, nil
}

func (r *orderRepository) FindAllOrdersWithExtraInfoByStoreId(storeId string) ([]*dto.OrderWithExtraInfoDTO ,error)  {
    query := `SELECT * FROM "public"."Order" WHERE "storeId" = $1;`
    rows, err := r.db.Query(query, storeId)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        logger.Error("Error while fetching Orders: ", zap.Error(err))
        return nil, err
    }
    defer rows.Close()

    var ordersWithExtraInfo []*dto.OrderWithExtraInfoDTO

    for rows.Next() {
        orderWithExtraInfo := &dto.OrderWithExtraInfoDTO{}

        err := rows.Scan(&orderWithExtraInfo.Id, &orderWithExtraInfo.StoreId,&orderWithExtraInfo.IsPaid ,&orderWithExtraInfo.Phone,&orderWithExtraInfo.Address, &orderWithExtraInfo.CreatedAt, &orderWithExtraInfo.UpdatedAt)
        
        if err != nil {
            logger.Error("Error scanning row:", zap.Error(err))
            return nil, err
        
        }
        ordersWithExtraInfo = append(ordersWithExtraInfo, orderWithExtraInfo)

    }

    if err := rows.Err(); err != nil {
        logger.Error("Error iterating rows:", zap.Error(err))
        return nil, err
    }
    
    return ordersWithExtraInfo, nil
}

func (r *orderRepository) Create(store *entity.Order) error {
    return nil
}

func (r *orderRepository) Update(store *entity.Order) (error) {
    return nil
}

func (r *orderRepository) Delete(id string) error {
    return nil
}

