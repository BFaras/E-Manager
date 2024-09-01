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

func (s *ColorService) GetAllColors() ([]*entity.Color, error) {
    colors, err := s.repository.FindAllColors()
    if err != nil {
        return nil, err
    }
    return colors, nil
}

func (s *ColorService) DeleteColor(id string) (error) {
    err := s.repository.Delete(id)
    if err != nil {
        return err
    }
    return nil
}

func (s *ColorService) UpdateColor(color *entity.Color) (error) {
    err := s.repository.Update(color)
    if err != nil {
        return err
    }
    return nil
}

func (s *ColorService) CreateColor(color *entity.Color) (error) {
    err := s.repository.Create(color)
    if err != nil {
        return err
    }
    return  nil
}


