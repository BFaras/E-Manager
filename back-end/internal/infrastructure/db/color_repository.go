package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"database/sql"
)

type colorRepository struct {
    db *sql.DB
}

func NewColorRepository(db *sql.DB) repository.ColorRepository {
    return &colorRepository{db: db}
}

func (r *colorRepository) FindByID(id string) (*entity.Color ,error) {
    color := &entity.Color{}
    query := `SELECT * FROM "public"."Color"stores WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(&color.ID, &color.StoreID, &color.Store,&color.Products,&color.Name,&color.Value,
        &color.CreatedAt, &color.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return color, nil
}

func (r *colorRepository) Create(store *entity.Color) error {
    return nil
}

func (r *colorRepository) Update(store *entity.Color) (*entity.Color, error) {
    return nil, nil
}

func (r *colorRepository) Delete(id string) error {
    return nil
}

