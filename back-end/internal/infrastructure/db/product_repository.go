package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"database/sql"
)

type productRepository struct {
    db *sql.DB
}

func NewProductRepository(db *sql.DB) repository.ProductRepository {
    return &productRepository{db: db}
}

func (r *productRepository) FindByID(id string) (*entity.Product ,error) {
    product := &entity.Product{}
    query := `SELECT * FROM "public"."Product"stores WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(&product.ID, &product.StoreID, &product.Store, &product.CategoryID,&product.Category,&product.Name,&product.Price,
        &product.IsFeatured,&product.IsArchived,&product.SizeID,&product.Size,&product.ColorID,&product.Color, &product.Images, &product.OrderItems, &product.CreatedAt, &product.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return product, nil
}

func (r *productRepository) Create(store *entity.Product) error {
    return nil
}

func (r *productRepository) Update(store *entity.Product) (*entity.Product, error) {
    return nil, nil
}

func (r *productRepository) Delete(id string) error {
    return nil
}

