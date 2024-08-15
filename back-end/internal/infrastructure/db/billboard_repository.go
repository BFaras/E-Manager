package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"database/sql"
)

type billboardRepository struct {
    db *sql.DB
}

func NewBillboardRepository(db *sql.DB) repository.BillboardRepository {
    return &billboardRepository{db: db}
}

func (r *billboardRepository) FindByID(id string) (*entity.Billboard ,error) {
    billboard := &entity.Billboard{}
    query := `SELECT * FROM "public"."Billboard"stores WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(&billboard.Id, &billboard.StoreId, &billboard.Label,&billboard.ImageUrl,
        &billboard.CreatedAt, &billboard.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return billboard, nil
}

func (r *billboardRepository) Create(store *entity.Billboard) error {
    return nil
}

func (r *billboardRepository) Update(store *entity.Billboard) (*entity.Billboard, error) {
    return nil, nil
}

func (r *billboardRepository) Delete(id string) error {
    return nil
}

