package db

import (
	"back-end/internal/domain/entity/dto"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/logger"
	"database/sql"
	"time"

	"go.uber.org/zap"
)

type DashboardInfoRepository struct {
	db *sql.DB
}

func NewDashboardInfoRepository(db *sql.DB) repository.DashboardInfoRepository {
    return &DashboardInfoRepository{db: db}
}

func (s *DashboardInfoRepository) FindTotalRevenue(storeID string) (float64, error) {
    query := `
        SELECT COALESCE(SUM(p.price), 0)
        FROM "public"."Order" o
        INNER JOIN "public"."OrderItem" oi ON o."id" = oi."orderId"
        INNER JOIN "public"."Product" p ON oi."productId" = p."id"
        WHERE o."storeId" = $1 
        AND o."isPaid" = TRUE;
    `
    var totalRevenue float64
    err := s.db.QueryRow(query, storeID).Scan(&totalRevenue)
    if err != nil {
        logger.Error("Could not calculate total revenue", zap.Error(err))
        return 0, err
    }
    return totalRevenue, nil
}

func (s *DashboardInfoRepository) FindTotalSales(storeID string) (int64, error) {
    query := `
        SELECT COUNT(*) FROM "public"."Order" o 
        WHERE o."isPaid" = TRUE AND o."storeId" = $1
    `
    var totalSales int64
    err := s.db.QueryRow(query, storeID).Scan(&totalSales)
    if err != nil {
        logger.Error("Could not calculate total sales", zap.Error(err))
        return 0, err
    }
    return totalSales, nil
}

func (s *DashboardInfoRepository) FindGraphRevenue(storeID string) ([]*dto.GraphDataDTO, error) {
	query := `
		SELECT o."createdAt", p."price"
		FROM "Order" o
		INNER JOIN "OrderItem" oi ON o."id" = oi."orderId"
		INNER JOIN "Product" p ON oi."productId" = p."id"
		WHERE o."storeId" = $1 AND o."isPaid" = TRUE
	`

	rows, err := s.db.Query(query, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	monthlyRevenue := make(map[int]float64)

	for rows.Next() {
		var createdAt time.Time
		var price float64

		err := rows.Scan(&createdAt, &price)
		if err != nil {
			return nil, err
		}

		month := int(createdAt.Month()) - 1
		monthlyRevenue[month] += price
		
	}

	graphData := []*dto.GraphDataDTO{
		{Name: "Jan", Total: 0},
		{Name: "Feb", Total: 0},
		{Name: "Mar", Total: 0},
		{Name: "Apr", Total: 0},
		{Name: "May", Total: 0},
		{Name: "Jun", Total: 0},
		{Name: "Jul", Total: 0},
		{Name: "Aug", Total: 0},
		{Name: "Sep", Total: 0},
		{Name: "Oct", Total: 0},
		{Name: "Nov", Total: 0},
		{Name: "Dec", Total: 0},
	}

	for month, revenue := range monthlyRevenue {
		graphData[int(month)].Total = revenue
	}

	return graphData, nil
}