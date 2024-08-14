package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"database/sql"
)

type orderItemRepository struct {
    db *sql.DB
}

func NewOrderItemRepository(db *sql.DB) repository.OrderItemRepository {
    return &orderItemRepository{db: db}
}

func (r *orderItemRepository) FindByID(id string) (*entity.OrderItem ,error) {
    orderItem := &entity.OrderItem{}
    query := `SELECT * FROM "public"."OrderItem"stores WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.Order,&orderItem.ProductID,&orderItem.Product)
    if err != nil {
        return nil, err
    }
    return orderItem, nil
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

