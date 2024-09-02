package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/logger"
	"database/sql"
    "go.uber.org/zap"
)

type ColorRepository struct {
    db *sql.DB
}

func NewColorRepository(db *sql.DB) repository.ColorRepository {
    return &ColorRepository{db: db}
}

func (r *ColorRepository) FindById(id string) (*entity.Color ,error) {
    color := &entity.Color{}
    query := `SELECT * FROM "public"."Color" WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(&color.Id, &color.StoreId, &color.Name,&color.Value,
        &color.CreatedAt, &color.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        logger.Error("Error while fetching color : ", zap.Error(err))
        return nil, err
    }
    return color, nil
}

func (r *ColorRepository) Delete(id string) error {
    query := `DELETE FROM "public"."Color" WHERE id = $1;`
    result, err := r.db.Exec(query, id)
    if err != nil {
        logger.Error("Error while deleting color : ",zap.Error(err))
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        logger.Error("Error : ",zap.Error(err))
        return err
    }

    if rowsAffected == 0 {
        logger.Error("no row found for : ",zap.String("colorId",id))
    }

    return nil
}

func (r *ColorRepository) FindAllColors() ([]*entity.Color, error) {
    var sizes []*entity.Color
    query := `SELECT * FROM "public"."Color"`
    rows, err := r.db.Query(query)
    if err != nil {
        logger.Error("Error while fetching all Colors : ", zap.Error(err))
        return nil, err
    }

    defer rows.Close()

    for rows.Next() {
        color := &entity.Color{}
        err := rows.Scan(&color.Id, &color.StoreId, &color.Name, &color.Value, &color.CreatedAt, &color.UpdatedAt)
        if err != nil {
            return nil, err
        }
        sizes = append(sizes, color)
    }
    return sizes, nil
}


func (r *ColorRepository) Create(color *entity.Color) error {
    query := `
        INSERT INTO "public"."Color" ("id", "storeId", "name","value","createdAt", "updatedAt")
        VALUES ($1, $2, $3, $4, $5, $6)
    `
    _, err := r.db.Exec(query, color.Id, color.StoreId, color.Name, color.Value, color.CreatedAt, color.UpdatedAt)
    if err != nil {
        logger.Error("Error : ",zap.Error(err))
        return err
    }

    return nil
}

func (r *ColorRepository) Update(color *entity.Color) (error) {
    query := `
        UPDATE "public"."Color"
        SET "storeId" = $1, "name" = $2, "value" = $3, "createdAt" = $4, "updatedAt" = $5
        WHERE "id" = $6
    `
    _, err := r.db.Exec(query, color.StoreId, color.Name, color.Value, color.CreatedAt, color.UpdatedAt, color.Id)
    if err != nil {
        logger.Error("Error: ", zap.Error(err))
        return err
    }

    return nil
}