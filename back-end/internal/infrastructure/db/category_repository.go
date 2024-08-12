package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"database/sql"
)

type categoryRepository struct {
    db *sql.DB
}

func NewCategoryRepository(db *sql.DB) repository.CategoryRepository {
    return &categoryRepository{db: db}
}

func (r *categoryRepository) FindByID(id string) (*entity.Category, error) {
    category := &entity.Category{}
    query := `SELECT * FROM "public"."Category" stores WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(&store.ID, &store.Name, &store.UserID, &store.CreatedAt, &store.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return store, nil

}

func (r *categoryRepository) Save(id string) (*entity.Category, error) {
    category := &entity.Category{}
    query := `SELECT * FROM "public"."Category" stores WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(&store.ID, &store.Name, &store.UserID, &store.CreatedAt, &store.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return store, nil

}