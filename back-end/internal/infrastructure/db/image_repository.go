package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/logger"
	"database/sql"

	"go.uber.org/zap"
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
        logger.Error("Error while finding an image: ",zap.Error(err))
        return nil, err
    }
    return image, nil
}

func (r *imageRepository) Create(image *entity.Image) error {
    query := `
    INSERT INTO "public"."Image" ("id", "productId", "url","createdAt", "updatedAt")
    VALUES ($1, $2, $3, $4, $5)
`
    _, err := r.db.Exec(query, image.Id, image.ProductId, image.URL, image.CreatedAt, image.UpdatedAt)
    if err != nil {
        logger.Error("Error while creating an image: ",zap.Error(err))
        return err
    }

return nil
}

func (r *imageRepository) Update(image *entity.Image) (*entity.Image, error) {
    return nil, nil
}

func (r *imageRepository) Delete(id string) error {
    return nil
}

