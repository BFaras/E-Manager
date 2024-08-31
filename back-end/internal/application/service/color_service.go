package service

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/db"
	"database/sql"
)

type ColorService struct {
    repository repository.ColorRepository
}

func NewColorService(database *sql.DB) *ColorService {
    return &ColorService{
        repository: db.NewColorRepository(database),
    }
}

func (s *ColorService) GetColor(id string) (*entity.Color, error) {
    color, err := s.repository.FindByID(id)
    if err != nil {
        return nil, err
    }
    return color, nil
}