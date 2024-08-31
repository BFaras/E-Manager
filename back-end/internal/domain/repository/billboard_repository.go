package repository

import ("back-end/internal/domain/entity")

type BillboardRepository interface {
    FindByID(id string) (*entity.Billboard, error)
    Create(billboard *entity.Billboard) error
    Update(billboard *entity.Billboard) (error)
    Delete(id string) error
    FindBillboardsByStoreId(storeId string) ([]*entity.Billboard, error) 
    FindActiveBillboard(storeId string) (*entity.Billboard, error)
}
