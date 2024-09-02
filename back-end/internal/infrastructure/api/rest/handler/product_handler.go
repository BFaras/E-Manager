package handler

import (
	"back-end/internal/domain/entity"
	"back-end/internal/infrastructure/logger"
	"database/sql"
	"net/http"
	"time"
    "back-end/internal/domain/entity/dto"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h* Handler) GetProductById(c echo.Context) (error) {
	logger.Debug("Fetching product by id...")
    productId := c.Param("productId")
    product, err := h.productService.GetProduct(productId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, product)
}

func (h *Handler) GetAllProductsWithExtraInformationByStoreId(c echo.Context) error {
    logger.Debug("Fetch products with extra category, size and colors by storeId...")
    storeId := c.Param("storeId")
    products, err := h.productService.GetAllProductsWithExtraInformationByStoreId(storeId)
    if err != nil {
        logger.Error("Error while trying to get all Products with extra info")
        return c.JSON(http.StatusInternalServerError, zap.Error(err))
    }
    return c.JSON(http.StatusOK, products)
}

func (h *Handler) GetAllProductsWithImageById(c echo.Context) error {
    logger.Debug("Fetch products with image by storeId...")
    productId := c.Param("productId")
    product, err := h.productService.GetAllProductsWithImageById(productId)
    logger.Debug("here is what it founds", zap.Reflect("product",product))
    if err != nil {
        logger.Error("Error while trying to get all Products with images")
        return c.JSON(http.StatusInternalServerError, zap.Error(err))
    }
    return c.JSON(http.StatusOK, product)
}

func (h *Handler) AddProduct(c echo.Context) error {
    logger.Debug("Adding new product...")

    storeId := c.Param("storeId")
	userId, ok := c.Get("userID").(string)
	if !ok || userId == "" {
        logger.Error("User ID missing or invalid")
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }
    var req *entity.Product
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    if !h.storeService.IsOwnerOfStore(storeId, userId) { 
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
        SizeId:     req.SizeId,
        ColorId:    req.ColorId,
        CreatedAt:  time.Now(),
        UpdatedAt:  time.Now(),
        Count: req.Count,
        IsDeleted: false,
    }

    err := h.productService.CreateProduct(product)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.NoContent(http.StatusOK)
}

func (h *Handler) UpdateProduct(c echo.Context) error {
    logger.Debug("Updating product...")

    productId := c.Param("productId")
    userId, ok := c.Get("userID").(string)
    if !ok || userId == "" {
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }

    var req dto.ProductWithExtraInfoDTO
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    product, err := h.productService.GetProduct(productId)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, err.Error())
        }
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    if !h.storeService.IsOwnerOfStore(product.StoreId, userId) { 
        return c.JSON(http.StatusForbidden, "You are not authorized to update this product")
    }

    product.Name = req.Name
    product.Price = req.Price
    product.CategoryId = req.CategoryId
    product.SizeId = req.SizeId
    product.ColorId = req.ColorId
    product.IsFeatured = req.IsFeatured
    product.IsArchived = req.IsArchived
    product.UpdatedAt = time.Now()

    if err := h.productService.UpdateProduct(product); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.NoContent(http.StatusOK)
}

func (h *Handler) DeleteProduct(c echo.Context) error {
    logger.Debug("Deleting product by id...")
    productId := c.Param("productId")
    storeId := c.Param("storeId")
    userId, ok := c.Get("userID").(string)
    if !ok || userId == "" {
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }
    if !h.storeService.IsOwnerOfStore(storeId, userId) { 
        return c.JSON(http.StatusForbidden, "You are not authorized to delete this product")
    }
    err := h.productService.DeleteProduct(productId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.NoContent(http.StatusOK)
}