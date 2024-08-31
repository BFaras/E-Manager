package handler

import (
	"net/http"
    "back-end/internal/infrastructure/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
    "back-end/internal/domain/entity"
    "time"
    "github.com/google/uuid"
    "database/sql"
)

type CategoryRequest struct {
    Name   string `json:"name"`
    BillboardId     string `json:"billboardId"`
}


func (h* Handler) GetCategoryById(c echo.Context) (error) {
	logger.Debug("Fetching category by id...")
    categoryId := c.Param("categoryId")
    category, err := h.categoryService.GetCategory(categoryId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, category)
}

func (h *Handler) GetCategoriesWithBillboard(c echo.Context) error {
    logger.Debug("Get categories with billboard by storeId...")
    storeId := c.Param("storeId")

    categories, err := h.categoryService.GetCategoriesWithBillboard(storeId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, zap.Error(err))
    }

    return c.JSON(http.StatusOK, categories)
}


func (h *Handler) DeleteCategory(c echo.Context) (error) {
	logger.Debug("Deleting category by id...")
	billboardId := c.Param("categoryId")
	storeId := c.Param("storeId")
	userID, ok := c.Get("userID").(string)
	if !ok || userID == "" {
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }
	if !h.storeService.IsOwnerOfStore(storeId, userID) { 
        return c.JSON(http.StatusForbidden, "You are not authorized to delete this category")
    }
	err := h.categoryService.DeleteCategory(billboardId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h *Handler) AddCategory(c echo.Context) error {
    logger.Debug("Adding new category...")

    storeId := c.Param("storeId")
	userID, ok := c.Get("userID").(string)
	if !ok || userID == "" {
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }
    var req CategoryRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

	if !h.storeService.IsOwnerOfStore(storeId, userID) { 
        return c.JSON(http.StatusForbidden, "You are not authorized to update this category")
    }

    category := &entity.Category{
        Id:        uuid.New().String(),
        StoreId:   storeId,
        Name:     req.Name,
        BillboardId:  req.BillboardId,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    err := h.categoryService.CreateCategory(category)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to create category")
    }

    return c.NoContent(http.StatusOK)
}

func (h *Handler) UpdateCategory(c echo.Context) error {
    logger.Debug("Updating category...")

    categoryId := c.Param("categoryId")
    userID, ok := c.Get("userID").(string)
    if !ok || userID == "" {
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }

    var req CategoryRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    category, err := h.categoryService.GetCategory(categoryId)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, "Category not found")
        }
        return c.JSON(http.StatusInternalServerError, "Failed to fetch category")
    }

    if !h.storeService.IsOwnerOfStore(category.StoreId, userID) { 
        return c.JSON(http.StatusForbidden, "You are not authorized to update this category")
    }

    category.Name = req.Name
    category.BillboardId = req.BillboardId
    category.UpdatedAt = time.Now()

    if err := h.categoryService.UpdateCategory(category); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.NoContent(http.StatusOK)
}