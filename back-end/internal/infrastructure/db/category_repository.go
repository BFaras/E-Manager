package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/entity/dto"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/logger"
	"database/sql"
    "go.uber.org/zap"
)

type CategoryRepository struct {
    db *sql.DB
}

func NewCategoryRepository(db *sql.DB) repository.CategoryRepository {
    return &CategoryRepository{db: db}
}

func (r *CategoryRepository) FindByID(id string) (*entity.Category, error) {
    category := &entity.Category{}
    query := `SELECT * FROM "public"."Category" stores WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(&category.Id, &category.StoreId, &category.BillboardId,
         &category.Name,&category.CreatedAt,&category.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return category, nil
}

func (r *CategoryRepository) FindCategoriesWithBillboard(storedId string) ([]*dto.CategoryWithBillboardDTO,error) {

    query := `SELECT 
    c."id", 
    c."storeId", 
    c."billboardId", 
    c."name", 
    c."createdAt", 
    c."updatedAt",
    b."id" AS "billboardId", 
    b."storeId" AS "billboardStoreId", 
    b."imageUrl", 
    b."createdAt" AS "billboardCreatedAt", 
    b."updatedAt" AS "billboardUpdatedAt"
    FROM "public"."Category" c
    LEFT JOIN "public"."Billboard" b
    ON c."billboardId" = b."id"
    ORDER BY c."createdAt" DESC;`

    rows,err := r.db.Query(query)

    if err != nil {
        logger.Error("Error executing query:", zap.Error(err))
        return nil, err
    }

    defer rows.Close()

    var categoriesWithBillboard []*dto.CategoryWithBillboardDTO

    for rows.Next() {
        var category *dto.CategoryWithBillboardDTO
        var billboard *entity.Billboard
        
        err := rows.Scan(
            &category.Id,
            &category.StoreId,
            &category.BillboardId,
            &category.Name,
            &category.CreatedAt,
            &category.UpdatedAt,
            &billboard.Id,
            &billboard.StoreId,
            &billboard.Label,
            &billboard.ImageUrl,
            &billboard.CreatedAt,
            &billboard.UpdatedAt,
            &billboard.IsActive,
        )
        if err != nil {
            logger.Error("Error scanning row:", zap.Error(err))
            return nil, err
        }
        
        category.Billboard = billboard
        categoriesWithBillboard = append(categoriesWithBillboard, category)
    }
    
    if err := rows.Err(); err != nil {
        logger.Error("Error iterating rows:", zap.Error(err))
        return nil, err
    }
    
    return categoriesWithBillboard, nil
}


func (r *CategoryRepository) Delete(id string) error {
    query := `DELETE FROM "public"."Category" WHERE id = $1;`
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

func (r *CategoryRepository) Create(billboard *entity.Category) error {
    query := `
        INSERT INTO "public"."Category" ("id", "storeId", "name","billboardId","createdAt", "updatedAt")
        VALUES ($1, $2, $3, $4, $5, $6)
    `
    _, err := r.db.Exec(query, billboard.Id, billboard.StoreId, billboard.Name, billboard.BillboardId, billboard.CreatedAt, billboard.UpdatedAt)
    if err != nil {
        logger.Error("Error : ",zap.Error(err))
        return err
    }

    return nil
}

func (r *CategoryRepository) Update(category *entity.Category) (error) {
    query := `
        UPDATE "public"."Category"
        SET "storeId" = $1, "name" = $2, "billboardId" = $3, "createdAt" = $4, "updatedAt" = $5
        WHERE "id" = $6
    `
    _, err := r.db.Exec(query, category.StoreId, category.Name, category.BillboardId, category.CreatedAt, category.UpdatedAt, category.Id)
    if err != nil {
        logger.Error("Error: ", zap.Error(err))
        return err
    }

    return nil
}