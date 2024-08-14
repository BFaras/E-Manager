package repository

import ("back-end/internal/domain/entity")

type BillboardRepository interface {
    FindByID(id string) (*entity.Billboard, error)
    Create(store *entity.Billboard) error
    Update(store *entity.Billboard) (*entity.Billboard, error)
    Delete(id string) error
}
