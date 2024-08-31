package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"database/sql"
)

type sizeRepository struct {
    db *sql.DB
}

func NewSizeRepository(db *sql.DB) repository.SizeRepository {
    return &sizeRepository{db: db}
}

func (r *sizeRepository) FindByID(id string) (*entity.Size ,error) {
    size:= &entity.Size{}
    query := `SELECT * FROM "public"."Size"stores WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(&size.Id, &size.StoreId,&size.Name ,&size.Value, &size.CreatedAt, &size.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return size, nil
}

func (r *sizeRepository) Create(store *entity.Size) error {
    return nil
}

func (r *sizeRepository) Update(store *entity.Size) (error) {
    return nil
}

func (r *sizeRepository) Delete(id string) error {
    return nil
}

