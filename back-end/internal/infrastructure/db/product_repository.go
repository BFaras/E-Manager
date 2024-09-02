package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/logger"
	"database/sql"
    "back-end/internal/domain/entity/dto"

	"go.uber.org/zap"
)

type productRepository struct {
    db *sql.DB
}

func NewProductRepository(db *sql.DB) repository.ProductRepository {
    return &productRepository{db: db}
}

func (r *productRepository) FindById(id string) (*entity.Product ,error) {
    product := &entity.Product{}
    query := `SELECT * FROM "public"."Product" WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(&product.Id, &product.StoreId,  &product.CategoryId,&product.Name,&product.Price,
        &product.IsFeatured,&product.IsArchived,&product.SizeId,&product.ColorId,  &product.CreatedAt, &product.UpdatedAt,&product.Count,&product.IsDeleted,)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        logger.Error("Error while fetching Product : ", zap.Error(err))
        return nil, err
    }
    return product, nil
}

func (r *productRepository) FindAllProductsWithExtraInformationByStoreId(storeId string) ([]*dto.ProductWithExtraInfoDTO, error) {
    var products []*dto.ProductWithExtraInfoDTO
    query := `SELECT 
        product.*, color.* , size.*, category.*
        FROM "public"."Product" product
        LEFT JOIN "public"."Color" color ON product."colorId" = color."id"
        LEFT JOIN "public"."Size" size ON product."sizeId" = size."id"
        LEFT JOIN "public"."Category" category ON product."categoryId" = category."id"
        WHERE product."storeId" = $1
        ORDER BY product."createdAt" DESC`

    rows, err := r.db.Query(query,storeId);
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        logger.Error("Error while fetching Products with extra Info : ", zap.Error(err))
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        product := &dto.ProductWithExtraInfoDTO{
            Color: &entity.Color{},
            Size: &entity.Size{},
            Category: &entity.Category{},
        }
    
        err := rows.Scan(
            &product.Id, &product.StoreId, &product.CategoryId, &product.Name, &product.Price,
            &product.IsFeatured, &product.IsArchived, &product.SizeId, &product.ColorId, &product.CreatedAt, &product.UpdatedAt,
            &product.Count,&product.IsDeleted,
            &product.Color.Id, &product.Color.StoreId, &product.Color.Name, &product.Color.Value, &product.Color.CreatedAt, &product.Color.UpdatedAt,
            &product.Size.Id, &product.Size.StoreId, &product.Size.Name, &product.Size.Value, &product.Size.CreatedAt, &product.Size.UpdatedAt,
            &product.Category.Id, &product.Category.StoreId, &product.Category.BillboardId, &product.Category.Name, &product.Category.CreatedAt, &product.Category.UpdatedAt,
        )
        if err != nil {
            logger.Error("Error while Scanning all Products With Extra Information: ", zap.Error(err))
            return nil, err
        }
        if (product.IsDeleted) {
            continue
        }
        products = append(products, product)
    }
    return products, nil
}

func (r *productRepository) FindAllProductsWithImageById(id string) (*dto.ProductWithImageDTO, error) {
    query := `SELECT product.*, image.*
              FROM "public"."Product" product
              LEFT JOIN "public"."Image" image
              ON product."id" = image."productId"
              WHERE product."id" = $1;`

    rows, err := r.db.Query(query, id)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        logger.Error("Error while fetching Products with Images: ", zap.String("id", id), zap.Error(err))
        return nil, err
    }
    defer rows.Close()

    product := &dto.ProductWithImageDTO{
        Images: []*entity.Image{},
    }

    hasRows := false
    for rows.Next() {
        hasRows = true 
        var image entity.Image
        err := rows.Scan(&product.Id, &product.StoreId, &product.CategoryId, &product.Name, &product.Price,
            &product.IsFeatured, &product.IsArchived, &product.SizeId, &product.ColorId, &product.CreatedAt, &product.UpdatedAt, &product.Count, &product.IsDeleted,
            &image.Id, &image.ProductId, &image.URL, &image.CreatedAt, &image.UpdatedAt)

        if err != nil {
            logger.Error("Error while scanning rows: ", zap.String("id", id), zap.Error(err))
            return nil, err
        }

        product.Images = append(product.Images, &image)
    }

    if err := rows.Err(); err != nil {
        logger.Error("Error while iterating rows: ", zap.String("id", id), zap.Error(err))
        return nil, err
    }

    if !hasRows {
        return nil, nil
    }

    return product, nil
}

func (r *productRepository) Create(product *dto.ProductWithImageDTO) error {
    query := `
        INSERT INTO "public"."Product" ("id", "storeId", "categoryId","name","price","isFeatured","isArchived","sizeId",
        "colorId","createdAt", "updatedAt","count", "isDeleted")
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
    `
    _, err := r.db.Exec(query,product.Id, product.StoreId,  product.CategoryId,product.Name,product.Price,
        product.IsFeatured,product.IsArchived,product.SizeId,product.ColorId, product.CreatedAt, product.UpdatedAt,product.Count,product.IsDeleted)
    if err != nil {
        logger.Error("Error while creating a product: ",zap.Error(err))
        return err
    }

    return nil
}

func (r *productRepository) Update(product *entity.Product) (error) {
    query := `
    UPDATE "public"."Product"
    SET "storeId" = $1, "categoryId" = $2, "name" = $3, "price" = $4, "isFeatured" = $5, "isArchived" = $6 ,"sizeId" = $7,
    "colorId" = $8 ,"createdAt" = $9 , "updatedAt" = $10, "count" = $11, "isDeleted" = $ 12
    WHERE "id" = $13
    `
    _, err := r.db.Exec(query, product.StoreId, product.CategoryId,product.Name,product.Price,
        product.IsFeatured,product.IsArchived,product.SizeId,product.ColorId, product.CreatedAt, product.UpdatedAt,product.Count,product.IsDeleted, product.Id)
    if err != nil {
        logger.Error("Error: ", zap.Error(err))
        return err
    }

    return nil
}

func (r *productRepository) Delete(id string) error {
    query := `Update "public"."Product"
    SET "isDeleted" = true
    WHERE "id" = $1`
    result, err := r.db.Exec(query, id)
    if err != nil {
        logger.Error("Error while deleting product : ",zap.Error(err))
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        logger.Error("Error : ",zap.Error(err))
        return err
    }

    if rowsAffected == 0 {
        logger.Error("no row found for : ",zap.String("productId",id))
    }

    return nil
}
