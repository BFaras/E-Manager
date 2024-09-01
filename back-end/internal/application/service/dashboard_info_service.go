package service

import (
	"back-end/internal/domain/entity/dto"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/db"
	"database/sql"
)

type DashboardInfoService struct {
	repository repository.DashboardInfoRepository
}

func NewDashboardInfoService(database *sql.DB) *DashboardInfoService {
    return &DashboardInfoService{
		repository: db.NewDashboardInfoRepository(database),
	}
}

func (s *DashboardInfoService) GetTotalRevenue(storeId string) (float64, error) {
    totalRevenue ,err := s.repository.FindTotalRevenue(storeId)
    if err != nil {
        return 0, err
    }
    return totalRevenue, nil
}

func (s *DashboardInfoService) GetTotalSales(storeId string) (int64, error) {
    totalSales, err := s.repository.FindTotalSales(storeId)
	if err != nil {
        return 0, err
    }
    return totalSales, nil
}

func (s *DashboardInfoService) GetGraphRevenue(storeId string) ([]*dto.GraphDataDTO, error) {
	totalGraphRevenue, err := s.repository.FindGraphRevenue(storeId)
	if err!= nil {
        return nil, err
    }
	return totalGraphRevenue, nil
}