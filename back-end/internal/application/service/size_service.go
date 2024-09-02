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

func (s *SizeService) GetSize(id string) (*entity.Size, error) {
    size, err := s.repository.FindById(id)
    if err != nil {
        return nil, err
    }
    return size, nil
}

func (s *SizeService) GetAllSizes() ([]*entity.Size, error) {
    size, err := s.repository.FindAllSizes()
    if err != nil {
        return nil, err
    }
    return size, nil
}

func (s *SizeService) DeleteSize(id string) (error) {
    err := s.repository.Delete(id)
    if err != nil {
        return err
    }
    return nil
}

func (s *SizeService) UpdateSize(size *entity.Size) (error) {
    err := s.repository.Update(size)
    if err != nil {
        return err
    }
    return nil
}

func (s *SizeService) CreateSize(size *entity.Size) (error) {
    err := s.repository.Create(size)
    if err != nil {
        return err
    }
    return  nil
}


