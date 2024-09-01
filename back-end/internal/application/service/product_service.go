package service

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/entity/dto"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/db"
	"database/sql"
)

type ProductService struct {
    repository repository.ProductRepository
}

func NewProductService(database *sql.DB) *ProductService {
    return &ProductService{
        repository: db.NewProductRepository(database),
    }
}

func (s *ProductService) GetProduct(id string) (*entity.Product, error) {
    product, err := s.repository.FindByID(id)
    if err != nil {
        return nil, err
    }
    return product, nil
}

func (s *ProductService) GetAllProductsWithExtraInformationByStoreId(storeId string) ([]*dto.ProductWithExtraInfoDTO, error) {
    products, err := s.repository.FindAllProductsWithExtraInformationByStoreId(storeId)
    if err != nil {
        return nil, err
    }
    return products, nil
}

func (s *ProductService) GetAllProductsWithImageById(storeId string) (*dto.ProductWithImageDTO, error) {
    products, err := s.repository.FindAllProductsWithImageById(storeId)
    if err != nil {
        return nil, err
    }
    return products, nil
}

func (s *ProductService) CreateProduct(product *entity.Product ) (error) {
     err := s.repository.Create(product)
    if err != nil {
        return err
    }
    return nil
}

func (s *ProductService) UpdateProduct(product *entity.Product ) (error) {
    err := s.repository.Update(product)
   if err != nil {
       return err
   }
   return nil
}

func (s *ProductService) DeleteProduct(id string ) (error) {
    err := s.repository.Delete(id)
   if err != nil {
       return err
   }
   return nil
}