package service

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/db"
	"database/sql"
)

type ImageService struct {
    repository repository.ImageRepository
}

func NewImageService(database *sql.DB) *ImageService {
    return &ImageService{
        repository: db.NewImageRepository(database),
    }
}

func (s *ImageService) GetImage(id string) (*entity.Image, error) {
    color, err := s.repository.FindByID(id)
    if err != nil {
        return nil, err
    }
    return color, nil
}