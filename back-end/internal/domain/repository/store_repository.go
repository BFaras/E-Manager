package repository

import ("back-end/internal/domain/entity")

type StoreRepository interface {
    FindById(id string) (*entity.Store, error)
    Create(store *entity.Store) error
    Update(store *entity.Store) (*entity.Store, error)
    Delete(id string) error
    FindByUserId(id string) (*entity.Store, error)
    
    // Add other methods as needed
}