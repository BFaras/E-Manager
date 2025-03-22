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

func (r *orderItemRepository) Create(item *entity.OrderItem) error {
	query := `
		INSERT INTO "public"."OrderItem" ("orderId", "productId")
		VALUES ($1, $2) RETURNING "id";
	`
	return r.db.QueryRow(query, item.OrderId, item.ProductId).Scan(&item.Id)
}

func (r *orderItemRepository) FindOrderItemsByOrderId(orderId string) ([]*entity.OrderItem, error) {
    var orderItems []*entity.OrderItem

    query := `SELECT "id", "orderId", "productId" FROM "public"."OrderItem" WHERE "orderId" = $1;`
    rows, err := r.db.Query(query, orderId)
    if err != nil {
        logger.Error("Error while fetching OrderItems by OrderId: ", zap.Error(err))
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        item := &entity.OrderItem{}
        if err := rows.Scan(&item.Id, &item.OrderId, &item.ProductId); err != nil {
            logger.Error("Error scanning OrderItem row: ", zap.Error(err))
            return nil, err
        }
        orderItems = append(orderItems, item)
    }

    if err := rows.Err(); err != nil {
        logger.Error("Error iterating OrderItem rows: ", zap.Error(err))
        return nil, err
    }

    return orderItems, nil
}
func (r *orderItemRepository) FindOrderItemByProductId(id string) (*entity.OrderItem, error) {
    orderItem := &entity.OrderItem{}
    query := `SELECT "id", "orderId", "productId" FROM "public"."OrderItem" WHERE "productId" = $1 LIMIT 1;`
    err := r.db.QueryRow(query, id).Scan(&orderItem.Id, &orderItem.OrderId, &orderItem.ProductId)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        logger.Error("Error while fetching OrderItem by ProductId: ", zap.Error(err))
        return nil, err
    }
    return orderItem, nil
}


func (r *orderItemRepository) Update(store *entity.OrderItem) (*entity.OrderItem, error) {
    return nil, nil
}

func (r *orderItemRepository) Delete(id string) error {
    return nil
}

