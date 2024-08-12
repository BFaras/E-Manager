package db

import ("back-end/internal/domain/entity")

type BillboardRepository interface {
    FindByID(id string) (*entity.Billboard, error)
    Save(billboard *entity.Billboard) error
    // Add other methods as needed
}
