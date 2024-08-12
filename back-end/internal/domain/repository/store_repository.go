package repository

import ("back-end/internal/domain/entity")

type StoreRepository interface {
    FindByID(id string) (*entity.Store, error)
    Save(store *entity.Store) error
    // Add other methods as needed
}