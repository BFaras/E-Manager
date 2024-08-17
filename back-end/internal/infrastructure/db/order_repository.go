package db

import (
	"back-end/internal/application/service"
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"database/sql"
)

type orderRepository struct {
    db *sql.DB
    DashboardInfoService *service.DashboardInfoService
}

func NewOrderRepository(db *sql.DB) repository.OrderRepository {
    return &orderRepository{db: db}
}

func (r *orderRepository) FindByID(id string) (*entity.Order ,error) {
    order := &entity.Order{}
    query := `SELECT * FROM "public"."Order" WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(&order.Id, &order.StoreId, &order.IsPaid,&order.Phone,
        &order.Address,&order.CreatedAt, &order.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return order, nil
}

func (r *orderRepository) CalculateRevenue(storeId string) (float64,error) {
    r.DashboardInfoService = service.NewDashboardInfoService(r.db)
    totalRevenue,err := r.DashboardInfoService.GetTotalRevenue(storeId)
    if err != nil {
        return 0, err
    }
    return totalRevenue, nil
}

func (r *orderRepository) CalculateSales(storeId string) (int64,error) {
    r.DashboardInfoService = service.NewDashboardInfoService(r.db)
    totalRevenue,err := r.DashboardInfoService.GetTotalSales(storeId)
    if err != nil {
        return 0, err
    }
    return totalRevenue, nil
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

