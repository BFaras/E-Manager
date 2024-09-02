package service

import (
    "database/sql"
    "back-end/internal/infrastructure/db"
    "back-end/internal/domain/repository"
    "back-end/internal/domain/entity"
)

type BillboardService struct {
    repository repository.BillboardRepository
}

func NewBilboardService(database *sql.DB) *BillboardService {
    return &BillboardService{
        repository: db.NewBillboardRepository(database),
    }
}

func (s *BillboardService) GetBillboard(id string) (*entity.Billboard, error) {
    billboard, err := s.repository.FindById(id)
    if err != nil {
        return nil, err
    }
    return billboard, nil
}

func (s *BillboardService) GetBillboardsByStoreId(id string) ([]*entity.Billboard, error) {
    billboards, err := s.repository.FindBillboardsByStoreId(id)
    if err != nil {
        return nil, err
    }
    return billboards, nil
}

func (s *BillboardService) GetActiveBillboard(id string) (*entity.Billboard, error) {
    billboard, err := s.repository.FindActiveBillboard(id)
    if err != nil {
        return nil, err
    }
    return billboard, nil
}

func (s *BillboardService) DeleteBillboard(id string) (error) {
    err := s.repository.Delete(id)
    if err != nil {
        return err
    }
    return nil
}


func (s *BillboardService) UpdateBillboard(billboard *entity.Billboard) (error) {
    err := s.repository.Update(billboard)
    if err != nil {
        return err
    }
    return nil
}


func (s *BillboardService) CreateBillboard(billboard *entity.Billboard) (error) {
    err := s.repository.Create(billboard)
    if err != nil {
        return err
    }
    return  nil
}

