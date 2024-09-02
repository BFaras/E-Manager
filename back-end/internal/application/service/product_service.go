package service

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/entity/dto"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/db"
	"database/sql"
    "time"
	"github.com/google/uuid"
)

type ProductService struct {
    prodcutRepository repository.ProductRepository
    imageRepository repository.ImageRepository
}

func NewProductService(database *sql.DB) *ProductService {
    return &ProductService{
        prodcutRepository: db.NewProductRepository(database),
        imageRepository: db.NewImageRepository(database),

    }
}

func (s *ProductService) GetProduct(id string) (*entity.Product, error) {
    product, err := s.prodcutRepository.FindById(id)
    if err != nil {
        return nil, err
    }
    return product, nil
}

func (s *ProductService) GetAllProductsWithExtraInformationByStoreId(storeId string) ([]*dto.ProductWithExtraInfoDTO, error) {
    products, err := s.prodcutRepository.FindAllProductsWithExtraInformationByStoreId(storeId)
    if err != nil {
        return nil, err
    }
    return products, nil
}

func (s *ProductService) GetAllProductsWithImageById(storeId string) (*dto.ProductWithImageDTO, error) {
    products, err := s.prodcutRepository.FindAllProductsWithImageById(storeId)
    if err != nil {
        return nil, err
    }
    return products, nil
}

func (s *ProductService) CreateProduct(product *dto.ProductWithImageDTO ) (error) {
     err := s.prodcutRepository.Create(product)
    if err != nil {
        return err
    }
    for _, image := range product.Images {
        image.Id = uuid.New().String()
        image.ProductId = product.Id
        image.CreatedAt = time.Now()
        image.UpdatedAt = time.Now()
        err := s.imageRepository.Create(image)
        if err != nil {
            return err
        }
    }
    return nil
}

func (s *ProductService) UpdateProduct(product *entity.Product ) (error) {
    err := s.prodcutRepository.Update(product)
   if err != nil {
       return err
   }
   return nil
}

func (s *ProductService) DeleteProduct(id string ) (error) {
    err := s.prodcutRepository.Delete(id)
   if err != nil {
       return err
   }
   return nil
}