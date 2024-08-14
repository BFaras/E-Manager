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
    err := r.db.QueryRow(query, id).Scan(&category.ID, &category.StoreID, &category.Store, &category.BillboardID, &category.Billboard,
        &category.Products, &category.Name,&category.CreatedAt,&category.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return category, nil

}

func (r *categoryRepository) Create(store *entity.Category) error {
    return nil
}

func (r *categoryRepository) Update(store *entity.Category) (*entity.Category, error) {
    return nil, nil
}

func (r *categoryRepository) Delete(id string) error {
    return nil
}

