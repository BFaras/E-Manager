package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/entity/dto"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/logger"
	"database/sql"
    "go.uber.org/zap"
)

type orderItemRepository struct {
    db *sql.DB
}

func NewOrderItemRepository(db *sql.DB) repository.OrderItemRepository {
    return &orderItemRepository{db: db}
}

func (r *orderItemRepository) FindById(id string) (*entity.OrderItem ,error) {
    orderItem := &entity.OrderItem{}
    query := `SELECT * FROM "public"."OrderItem"stores WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(&orderItem.Id, &orderItem.OrderId, &orderItem.ProductId)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        logger.Error("Error while fetching OrderItem: ",zap.Error(err))
        return nil, err
    }
    return orderItem, nil
}

func (r *orderItemRepository) FindAllOrderItemsByOrderId(orderId string) ([]*dto.OrderItemWithProductDTO ,error) {
    var orderItemsWithProduct []*dto.OrderItemWithProductDTO
    query := `SELECT "orderItem".* 
    FROM "public"."OrderItem" "orderItem"
    INNER JOIN "public"."Order" "order" 
    ON "orderItem"."orderId" = "order"."id"
    WHERE "orderItem"."orderId" = $1`
    rows, err := r.db.Query(query, orderId)
    if err!= nil {
        logger.Error("Error while fetching all OrderItems: ",zap.Error(err))
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        orderItemWithProduct := &dto.OrderItemWithProductDTO{}

        err := rows.Scan(&orderItemWithProduct.Id, &orderItemWithProduct.OrderId,&orderItemWithProduct.ProductId)
        
        if err != nil {
            logger.Error("Error scanning orderItem row:", zap.Error(err))
            return nil, err
        
        }
        orderItemsWithProduct = append(orderItemsWithProduct, orderItemWithProduct)

    }

    if err := rows.Err(); err != nil {
        logger.Error("Error iterating rows:", zap.Error(err))
        return nil, err
    }
    
    return orderItemsWithProduct, nil
}

func (r *orderItemRepository) Create(store *entity.OrderItem) error {
    return nil
}

func (r *orderItemRepository) Update(store *entity.OrderItem) (*entity.OrderItem, error) {
    return nil, nil
}

func (r *orderItemRepository) Delete(id string) error {
    return nil
}

