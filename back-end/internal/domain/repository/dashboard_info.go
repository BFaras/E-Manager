package repository

import (
	"back-end/internal/domain/entity/dto"
)

type DashboardInfoRepository interface {
	FindTotalRevenue(storeID string) (float64, error) 
	FindTotalSales(storeID string) (int64, error) 
	FindGraphRevenue(storeID string) ([]*dto.GraphData, error) 
}


