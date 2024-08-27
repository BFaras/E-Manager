package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/logger"
	"database/sql"
	"go.uber.org/zap"
)

type billboardRepository struct {
    db *sql.DB
}

func NewBillboardRepository(db *sql.DB) repository.BillboardRepository {
    return &billboardRepository{db: db}
}

func (r *billboardRepository) FindByID(id string) (*entity.Billboard, error) {
    billboard := &entity.Billboard{}
    query := `SELECT * FROM "public"."Billboard" WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(
        &billboard.Id,
        &billboard.StoreId,
        &billboard.Label,
        &billboard.ImageUrl,
        &billboard.CreatedAt,
        &billboard.UpdatedAt,
        &billboard.IsActive,
    )
    
    if err != nil {
        logger.Error("Error : ",zap.Error(err))
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }

    return billboard, nil
}

func (r *billboardRepository) GetBillboardsByStoreId(storeId string) ([]*entity.Billboard, error) {
	query := `
        SELECT *
        FROM "public"."Billboard"
        WHERE "storeId" = $1
        ORDER BY "createdAt" DESC
    `
	rows, err := r.db.Query(query, storeId)
	if err != nil {
        logger.Error("Error : ",zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var billboards []*entity.Billboard
	for rows.Next() {
		b := &entity.Billboard{}
		err := rows.Scan(&b.Id, &b.StoreId,&b.Label,&b.ImageUrl,&b.CreatedAt,&b.UpdatedAt,&b.IsActive)
		if err != nil {
            logger.Error("Error : ",zap.Error(err))
			return nil, err
		}
		billboards = append(billboards, b)
	}
	logger.Debug("Getting billboards for store: ", zap.Reflect("billboards",billboards));

	if err = rows.Err(); err != nil {
        logger.Error("Error : ",zap.Error(err))
 		return nil, err
	}

	return billboards, nil
}

func (r *billboardRepository) GetActiveBillboard(storeId string) (*entity.Billboard, error) {
	billboard := &entity.Billboard{}
	query := `
	SELECT *
	FROM "public"."Billboard"
	WHERE "storeId" = $1 AND "isActive" = true
	`
	err := r.db.QueryRow(query, storeId).Scan(&billboard.Id, &billboard.StoreId, &billboard.Label,&billboard.ImageUrl,
        &billboard.CreatedAt, &billboard.UpdatedAt,&billboard.IsActive)
    if err != nil {
        logger.Error("Error : ",zap.Error(err))
        return nil, err
    }
    return billboard, nil
}


func (r *billboardRepository) Delete(id string) error {
    query := `DELETE FROM "public"."Billboard" WHERE id = $1;`
    result, err := r.db.Exec(query, id)
    if err != nil {
        logger.Error("Error : ",zap.Error(err))
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        logger.Error("Error : ",zap.Error(err))
        return err
    }

    if rowsAffected == 0 {
        logger.Error("no row found for : ",zap.String("billboardId",id))
    }

    return nil
}

func (r *billboardRepository) Create(billboard *entity.Billboard) error {
    query := `
        INSERT INTO "public"."Billboard" ("id", "storeId", "label", "imageUrl", "createdAt", "updatedAt", "isActive")
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `
    _, err := r.db.Exec(query, billboard.Id, billboard.StoreId, billboard.Label, billboard.ImageUrl, billboard.CreatedAt, billboard.UpdatedAt, billboard.IsActive)
    if err != nil {
        logger.Error("Error : ",zap.Error(err))
        return err
    }

    return nil
}


func (r *billboardRepository) Update(store *entity.Billboard) (*entity.Billboard, error) {
    return nil, nil
}

