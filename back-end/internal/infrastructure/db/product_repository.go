package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/logger"
	"database/sql"
    "back-end/internal/domain/entity/dto"
    "time"
	"go.uber.org/zap"
    "fmt"
)

type productRepository struct {
    db *sql.DB
}

func NewProductRepository(db *sql.DB) repository.ProductRepository {
    return &productRepository{db: db}
}

func (r *productRepository) FindById(id string) (*entity.Product, error) {
    product := &entity.Product{}
    query := `SELECT "id", "storeId", "categoryId", "name", "price", "isFeatured",
                     "isArchived", "sizeId", "colorId", "createdAt", "updatedAt", "count", "isDeleted"
              FROM "public"."Product" WHERE "id" = $1;`

    err := r.db.QueryRow(query, id).Scan(
        &product.Id, &product.StoreId, &product.CategoryId, &product.Name, &product.Price,
        &product.IsFeatured, &product.IsArchived, &product.SizeId, &product.ColorId,
        &product.CreatedAt, &product.UpdatedAt, &product.Count, &product.IsDeleted,
    )
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        logger.Error("Error while fetching Product: ", zap.Error(err))
        return nil, err
    }
    return product, nil
}

func (r *productRepository) FindAllProductsWithExtraInformationByStoreId(storeId string, filter dto.ProductFilterDTO) ([]*dto.ProductWithExtraInfoDTO, error) {
    productMap := make(map[string]*dto.ProductWithExtraInfoDTO)

    baseQuery := `
        SELECT 
            product."id", product."storeId", product."categoryId", product."name", product."price",
            product."isFeatured", product."isArchived", product."sizeId", product."colorId", 
            product."createdAt", product."updatedAt", product."count", product."isDeleted",
            color."id", color."storeId", color."name", color."value", color."createdAt", color."updatedAt",
            size."id", size."storeId", size."name", size."value", size."createdAt", size."updatedAt",
            category."id", category."storeId", category."billboardId", category."name", category."createdAt", category."updatedAt",
            image."id", image."productId", image."url", image."createdAt", image."updatedAt"
        FROM "public"."Product" product
        LEFT JOIN "public"."Color" color ON product."colorId" = color."id"
        LEFT JOIN "public"."Size" size ON product."sizeId" = size."id"
        LEFT JOIN "public"."Category" category ON product."categoryId" = category."id"
        LEFT JOIN "public"."Image" image ON product."id" = image."productId"
        WHERE product."storeId" = $1
    `

    args := []interface{}{storeId}
    argIndex := 2
    if filter.ColorId != "" {
        baseQuery += fmt.Sprintf(` AND product."colorId" = $%d`, argIndex)
        args = append(args, filter.ColorId)
        argIndex++
    }
    if filter.SizeId != "" {
        baseQuery += fmt.Sprintf(` AND product."sizeId" = $%d`, argIndex)
        args = append(args, filter.SizeId)
        argIndex++
    }
    if filter.CategoryId != "" {
        baseQuery += fmt.Sprintf(` AND product."categoryId" = $%d`, argIndex)
        args = append(args, filter.CategoryId)
        argIndex++
    }
    if filter.IsFeatured == "true" || filter.IsFeatured == "false" {
        baseQuery += fmt.Sprintf(` AND product."isFeatured" = $%d`, argIndex)
        isFeaturedBool := (filter.IsFeatured == "true")
        args = append(args, isFeaturedBool)
        argIndex++
    }

    baseQuery += ` ORDER BY product."createdAt" DESC`

    rows, err := r.db.Query(baseQuery, args...)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        logger.Error("Error while fetching Products with filters: ", zap.Error(err))
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var (
            productId, pStoreId, categoryId, name string
            price                                   float64
            isFeatured, isArchived                  bool
            sizeId, colorId                       string
            createdAt, updatedAt                    time.Time
            count                                   int
            isDeleted                               bool

            colorIdDB, colorStoreId, colorName, colorValue string
            colorCreatedAt, colorUpdatedAt                   time.Time

            sizeIdDB, sizeStoreId, sizeName, sizeValue string
            sizeCreatedAt, sizeUpdatedAt                 time.Time

            categoryIdDB, categoryStoreId, billboardId, categoryName string
            categoryCreatedAt, categoryUpdatedAt                       time.Time

            imageId, imageProductId, imageURL sql.NullString
            imageCreatedAt, imageUpdatedAt      sql.NullTime
        )

        err := rows.Scan(
            &productId, &pStoreId, &categoryId, &name, &price,
            &isFeatured, &isArchived, &sizeId, &colorId,
            &createdAt, &updatedAt, &count, &isDeleted,
            &colorIdDB, &colorStoreId, &colorName, &colorValue, &colorCreatedAt, &colorUpdatedAt,
            &sizeIdDB, &sizeStoreId, &sizeName, &sizeValue, &sizeCreatedAt, &sizeUpdatedAt,
            &categoryIdDB, &categoryStoreId, &billboardId, &categoryName, &categoryCreatedAt, &categoryUpdatedAt,
            &imageId, &imageProductId, &imageURL, &imageCreatedAt, &imageUpdatedAt,
        )
        if err != nil {
            logger.Error("Error while scanning product row: ", zap.Error(err))
            return nil, err
        }

        if isDeleted {
            continue
        }

        prod, exists := productMap[productId]
        if !exists {

            prod = &dto.ProductWithExtraInfoDTO{
                Id:         productId,
                StoreId:    pStoreId,
                CategoryId: categoryId,
                Name:       name,
                Price:      price,
                IsFeatured: isFeatured,
                IsArchived: isArchived,
                SizeId:     sizeId,
                ColorId:    colorId,
                CreatedAt:  createdAt,
                UpdatedAt:  updatedAt,
                Count:      count,
                IsDeleted:  isDeleted,
                Color: &entity.Color{
                    Id:        colorIdDB,
                    StoreId:   colorStoreId,
                    Name:      colorName,
                    Value:     colorValue,
                    CreatedAt: colorCreatedAt,
                    UpdatedAt: colorUpdatedAt,
                },
                Size: &entity.Size{
                    Id:        sizeIdDB,
                    StoreId:   sizeStoreId,
                    Name:      sizeName,
                    Value:     sizeValue,
                    CreatedAt: sizeCreatedAt,
                    UpdatedAt: sizeUpdatedAt,
                },
                Category: &entity.Category{
                    Id:         categoryIdDB,
                    StoreId:    categoryStoreId,
                    BillboardId: billboardId,
                    Name:       categoryName,
                    CreatedAt:  categoryCreatedAt,
                    UpdatedAt:  categoryUpdatedAt,
                },
                Images: []*entity.Image{},
            }
            productMap[productId] = prod
        }

        if imageId.Valid {
            img := &entity.Image{
                Id:        imageId.String,
                ProductId: imageProductId.String,
                URL:       imageURL.String,
            }
            if imageCreatedAt.Valid {
                img.CreatedAt = imageCreatedAt.Time
            }
            if imageUpdatedAt.Valid {
                img.UpdatedAt = imageUpdatedAt.Time
            }
            prod.Images = append(prod.Images, img)
        }
    }

    if err := rows.Err(); err != nil {
        logger.Error("Error while iterating rows: ", zap.Error(err))
        return nil, err
    }

    products := make([]*dto.ProductWithExtraInfoDTO, 0, len(productMap))
    for _, p := range productMap {
        products = append(products, p)
    }

    return products, nil
}


func (r *productRepository) FindAllProductsWithImageById(id string) (*dto.ProductWithImageDTO, error) {
    query := `SELECT 
        product."id", product."storeId", product."categoryId", product."name", product."price",
        product."isFeatured", product."isArchived", product."sizeId", product."colorId",
        product."createdAt", product."updatedAt", product."count", product."isDeleted",
        image."id", image."productId", image."url", image."createdAt", image."updatedAt"
        FROM "public"."Product" product
        LEFT JOIN "public"."Image" image ON product."id" = image."productId"
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
        err := rows.Scan(
            &product.Id, &product.StoreId, &product.CategoryId, &product.Name, &product.Price,
            &product.IsFeatured, &product.IsArchived, &product.SizeId, &product.ColorId,
            &product.CreatedAt, &product.UpdatedAt, &product.Count, &product.IsDeleted,
            &image.Id, &image.ProductId, &image.URL, &image.CreatedAt, &image.UpdatedAt,
        )

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
    INSERT INTO "public"."Product" 
    ("id", "storeId", "categoryId", "name", "price", "count", "isFeatured", "isArchived", "isDeleted", "sizeId",
    "colorId", "createdAt", "updatedAt")
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
`
_, err := r.db.Exec(query,
    product.Id, product.StoreId, product.CategoryId, product.Name, product.Price, 
    product.Count, product.IsFeatured, product.IsArchived, product.IsDeleted,
    product.SizeId, product.ColorId, product.CreatedAt, product.UpdatedAt)
    if err != nil {
        logger.Error("Error while creating a product: ",zap.Error(err))
        return err
    }

    return nil
}

func (r *productRepository) Update(product *entity.Product) error {
    query := `
    UPDATE "public"."Product"
    SET "storeId" = $1, "categoryId" = $2, "name" = $3, "price" = $4, 
        "isFeatured" = $5, "isArchived" = $6, "sizeId" = $7, "colorId" = $8,
        "count" = $9, "isDeleted" = $10, "updatedAt" = $11
    WHERE "id" = $12
    `
    _, err := r.db.Exec(query, 
        product.StoreId, product.CategoryId, product.Name, product.Price,
        product.IsFeatured, product.IsArchived, product.SizeId, product.ColorId, 
        product.Count, product.IsDeleted, product.UpdatedAt, product.Id)

    if err != nil {
        logger.Error("Error updating product: ", zap.Error(err))
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
