package service

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/db"
	"database/sql"
)

type StoreService struct {
    repository repository.StoreRepository
}

func NewStoreService(database *sql.DB) *StoreService {
    return &StoreService{
        repository: db.NewStoreRepository(database),
    }
}

func (s *StoreService) GetStore(id string) (*entity.Store, error) {
    store, err := s.repository.FindById(id)
    if err != nil {
        return nil, err
    }
    return store, nil
}


func (s *StoreService) GetAllStores() ([]*entity.Store, error) {
    stores,err := s.repository.FindAllStores()
    if err != nil {
        return nil, err
    }
    return stores, nil
}

func (s *StoreService) GetByUserId(userId string) (*entity.Store ,error) {
    store,err := s.repository.FindByUserId(userId)
    if err != nil {
        return nil, err
    }
    return store, nil
}

func (s *StoreService) GetAllByUserId(userId string) ([]*entity.Store, error) {
    stores,err := s.repository.FindAllByUserId(userId)
    if err != nil {
        return nil, err
    }
    return stores, nil
}

func (s *StoreService) GetByIdAndUserId(id string ,userId string) (*entity.Store, error) {
    store, err := s.repository.FindByIdAndUserId(id,userId)
    if err!= nil {
        return nil, err
    }
    return store, nil
}

func (s *StoreService) IsOwnerOfStore(id string, userId string) (bool) {
    isOwnerOfStore := s.repository.IsOwnerOfStore(id,userId)
    return isOwnerOfStore
}

func (s *StoreService) CreateStore(store *entity.Store) error {
    return nil
}

func (s *StoreService) UpdateStore(store *entity.Store) (error) {
    return nil
}

func (s *StoreService) DeleteStore(id string) error {
    return nil
}



