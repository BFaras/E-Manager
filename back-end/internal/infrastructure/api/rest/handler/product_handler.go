package handler

import (
	"back-end/internal/domain/entity"
	"back-end/internal/infrastructure/logger"
	"database/sql"
	"net/http"
	"time"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type CreateProductRequest struct {
    CategoryId string  `json:"categoryId"`
    Name       string  `json:"name"`
    Price      float64 `json:"price"`
    IsFeatured bool    `json:"isFeatured"`
    IsArchived bool    `json:"isArchived"`
    SizeID     string  `json:"sizeId"`
    ColorID    string  `json:"colorId"`
}

type UpdateProductRequest struct {
    CategoryId string  `json:"categoryId"`
    Name       string  `json:"name"`
    Price      float64 `json:"price"`
    IsFeatured bool    `json:"isFeatured"`
    IsArchived bool    `json:"isArchived"`
    SizeID     string  `json:"sizeId"`
    ColorID    string  `json:"colorId"`
}


func (h *Handler) AddProduct(c echo.Context) error {
    logger.Debug("Adding new product...")

    storeId := c.Param("storeId")
	userId, ok := c.Get("userID").(string)
	if !ok || userId == "" {
        logger.Error("User ID missing or invalid")
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }
    var req CreateProductRequest
    if err := c.Bind(&req); err != nil {
        logger.Error("Failed to bind request: ", zap.Reflect("error", err))
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    if !h.storeService.IsOwnerOfStore(storeId, userId) { 
        logger.Error("User is not authorized to add a product", zap.String("userId", userId))
        return c.JSON(http.StatusForbidden, "You are not authorized to add this product")
    }

    product := &entity.Product{
        Id:         uuid.New().String(),
        StoreId:    storeId,
        CategoryId: req.CategoryId,
        Name:       req.Name,
        Price:      req.Price,
        IsFeatured: req.IsFeatured,
        IsArchived: req.IsArchived,
        SizeID:     req.SizeID,
        ColorID:    req.ColorID,
        CreatedAt:  time.Now(),
        UpdatedAt:  time.Now(),
    }

    err := h.productService.CreateProduct(product)
    if err != nil {
        logger.Error("Failed to create product: ", zap.Reflect("error", err))
        return c.JSON(http.StatusInternalServerError, "Failed to create product")
    }

    return c.NoContent(http.StatusOK)
}

func (h *Handler) UpdateProduct(c echo.Context) error {
    logger.Debug("Updating product...")

    productId := c.Param("productId")
    userId, ok := c.Get("userID").(string)
    if !ok || userId == "" {
        logger.Error("User ID missing or invalid")
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }

    var req UpdateProductRequest
    if err := c.Bind(&req); err != nil {
        logger.Error("Failed to bind request: ", zap.Error(err))
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    logger.Debug("This is the request: ", zap.Reflect("req", req))

    product, err := h.productService.GetProduct(productId)
    if err != nil {
        if err == sql.ErrNoRows {
            logger.Error("Product not found", zap.String("productId", productId))
            return c.JSON(http.StatusNotFound, "Product not found")
        }
        logger.Error("Failed to fetch product: ", zap.Error(err))
        return c.JSON(http.StatusInternalServerError, "Failed to fetch product")
    }

    if !h.storeService.IsOwnerOfStore(product.StoreId, userId) { 
        logger.Error("User is not authorized to update this product", zap.String("userId", userId), zap.String("productId", productId))
        return c.JSON(http.StatusForbidden, "You are not authorized to update this product")
    }

    product.Name = req.Name
    product.Price = req.Price
    product.CategoryId = req.CategoryId
    product.SizeID = req.SizeID
    product.ColorID = req.ColorID
    product.IsFeatured = req.IsFeatured
    product.IsArchived = req.IsArchived
    product.UpdatedAt = time.Now()

    logger.Debug("This is the updated product: ", zap.Reflect("product", product))

    if err := h.productService.UpdateProduct(product); err != nil {
        logger.Error("Failed to update product: ", zap.Error(err))
        return c.JSON(http.StatusInternalServerError, "Failed to update product")
    }

    return c.NoContent(http.StatusOK)
}

func (h *Handler) DeleteProduct(c echo.Context) error {
    logger.Debug("Deleting product by id...")
    productId := c.Param("productId")
    storeId := c.Param("storeId")
    userId, ok := c.Get("userID").(string)
    if !ok || userId == "" {
        logger.Error("User ID missing or invalid")
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }
    if !h.storeService.IsOwnerOfStore(storeId, userId) { 
        logger.Error("User is not authorized to delete this product", zap.String("userId", userId), zap.String("productId", productId))
        return c.JSON(http.StatusForbidden, "You are not authorized to delete this product")
    }
    err := h.productService.DeleteProduct(productId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.NoContent(http.StatusOK)
}