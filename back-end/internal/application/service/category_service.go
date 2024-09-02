package service

import (
    "database/sql"
	"back-end/internal/domain/entity/dto"
    "back-end/internal/infrastructure/db"
    "back-end/internal/domain/repository"
    "back-end/internal/domain/entity"
)

type CategoryService struct {
    repository repository.CategoryRepository
}

func NewCategoryService(database *sql.DB) *CategoryService {
    return &CategoryService{
        repository: db.NewCategoryRepository(database),
    }
}

func (s *CategoryService) GetCategory(id string) (*entity.Category, error) {
    categories, err := s.repository.FindById(id)
    if err != nil {
        return nil, err
    }
    return categories, nil
}

func (s *CategoryService) GetCategoriesWithBillboard(storeId string) ([]*dto.CategoryWithBillboardDTO, error) {
    categories, err := s.repository.FindCategoriesWithBillboard(storeId)
    if err != nil {
        return nil, err
    }
    return categories, nil
}

func (s *CategoryService) DeleteCategory(id string) (error) {
    err := s.repository.Delete(id)
    if err != nil {
        return err
    }
    return nil
}


func (s *CategoryService) UpdateCategory(category *entity.Category) (error) {
    err := s.repository.Update(category)
    if err != nil {
        return err
    }
    return nil
}


func (s *CategoryService) CreateCategory(category *entity.Category) (error) {
    err := s.repository.Create(category)
    if err != nil {
        return err
    }
    return  nil
}

