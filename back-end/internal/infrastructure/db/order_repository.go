package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"database/sql"
)

type orderRepository struct {
    db *sql.DB
}

func NewOrderRepository(db *sql.DB) repository.OrderRepository {
    return &orderRepository{db: db}
}

func (r *orderRepository) FindByID(id string) (*entity.Order ,error) {
    order := &entity.Order{}
    query := `SELECT * FROM "public"."Order"stores WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(&order.Id, &order.StoreId, &order.IsPaid,&order.Phone,
        &order.Address,&order.CreatedAt, &order.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return order, nil
}

func (r *orderRepository) Create(store *entity.Order) error {
    return nil
}

func (r *orderRepository) Update(store *entity.Order) (*entity.Order, error) {
    return nil, nil
}

func (r *orderRepository) Delete(id string) error {
    return nil
}

