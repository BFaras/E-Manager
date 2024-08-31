package service

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/db"
	"database/sql"
)

type SizeService struct {
    repository repository.SizeRepository
}

func NewSizeService(database *sql.DB) *SizeService {
    return &SizeService{
        repository: db.NewSizeRepository(database),
    }
}

func (s *SizeService) GetProduct(id string) (*entity.Size, error) {
    size, err := s.repository.FindByID(id)
    if err != nil {
        return nil, err
    }
    return size, nil
}