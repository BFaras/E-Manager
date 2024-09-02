package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/logger"
	"database/sql"
    "go.uber.org/zap"
)

type SizeRepository struct {
    db *sql.DB
}

func NewSizeRepository(db *sql.DB) repository.SizeRepository {
    return &SizeRepository{db: db}
}

func (r *SizeRepository) FindById(id string) (*entity.Size ,error) {
    size:= &entity.Size{}
    query := `SELECT * FROM "public"."Size" stores WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(&size.Id, &size.StoreId,&size.Name ,&size.Value, &size.CreatedAt, &size.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        logger.Error("Error while fetching Size : ", zap.Error(err))
        return nil, err
    }
    return size, nil
}

func (r *SizeRepository) FindAllSizes(storeId string) ([]*entity.Size, error) {
    var sizes []*entity.Size
    query := `SELECT * FROM "public"."Size" WHERE "storeId" = $1;`
    rows, err := r.db.Query(query, storeId)
    if err != nil {
        logger.Error("Error while fetching all Sizes : ", zap.Error(err))
        return nil, err
    }

    defer rows.Close()

    for rows.Next() {
        size := &entity.Size{}
        err := rows.Scan(&size.Id, &size.StoreId, &size.Name, &size.Value, &size.CreatedAt, &size.UpdatedAt)
        if err != nil {
            return nil, err
        }
        sizes = append(sizes, size)
    }
    return sizes, nil
}

func (r *SizeRepository) Delete(id string) error {
    query := `DELETE FROM "public"."Size" WHERE id = $1;`
    result, err := r.db.Exec(query, id)
    if err != nil {
        logger.Error("Error while deleting size : ",zap.Error(err))
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        logger.Error("Error : ",zap.Error(err))
        return err
    }

    if rowsAffected == 0 {
        logger.Error("no row found for : ",zap.String("sizeId",id))
    }

    return nil
}

func (r *SizeRepository) Create(size *entity.Size) error {
    query := `
        INSERT INTO "public"."Size" ("id", "storeId", "name","value","createdAt", "updatedAt")
        VALUES ($1, $2, $3, $4, $5, $6)
    `
    _, err := r.db.Exec(query, size.Id, size.StoreId, size.Name, size.Value, size.CreatedAt, size.UpdatedAt)
    if err != nil {
        logger.Error("Error : ",zap.Error(err))
        return err
    }

    return nil
}

func (r *SizeRepository) Update(size *entity.Size) (error) {
    query := `
        UPDATE "public"."Size"
        SET "storeId" = $1, "name" = $2, "value" = $3, "createdAt" = $4, "updatedAt" = $5
        WHERE "id" = $6
    `
    _, err := r.db.Exec(query, size.StoreId, size.Name, size.Value, size.CreatedAt, size.UpdatedAt, size.Id)
    if err != nil {
        logger.Error("Error: ", zap.Error(err))
        return err
    }

    return nil
}
