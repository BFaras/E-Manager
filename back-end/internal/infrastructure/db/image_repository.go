package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"database/sql"
)

type imageRepository struct {
    db *sql.DB
}

func NewImageRepository(db *sql.DB) repository.ImageRepository {
    return &imageRepository{db: db}
}

func (r *imageRepository) FindById(id string) (*entity.Image ,error) {
    image := &entity.Image{}
    query := `SELECT * FROM "public"."Image"stores WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(&image.Id, &image.ProductId, &image.URL, &image.CreatedAt, &image.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return image, nil
}

func (r *imageRepository) Create(store *entity.Image) error {
    return nil
}

func (r *imageRepository) Update(store *entity.Image) (*entity.Image, error) {
    return nil, nil
}

func (r *imageRepository) Delete(id string) error {
    return nil
}

