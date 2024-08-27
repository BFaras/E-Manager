package repository

import ("back-end/internal/domain/entity")

type StoreRepository interface {
    FindById(id string) (*entity.Store, error)
    FindAllStores() ([]*entity.Store, error)
    Create(store *entity.Store) error
    Update(store *entity.Store) (*entity.Store, error)
    Delete(id string) error
    FindByUserId(userId string) (*entity.Store, error)
    FindAllByUserId(userId string) ([]*entity.Store, error)
    FindByIdAndUserId(id string,userId string) (*entity.Store, error)
    IsOwnerOfStore(id string,userId string) (bool)
    
    
}