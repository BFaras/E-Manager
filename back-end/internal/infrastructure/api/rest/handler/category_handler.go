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

func (h *Handler) GetCategoriesWithBillboard(c echo.Context) error {
    storeId := c.Param("storeId")

    categories, err := h.categoryService.GetCategoriesWithBillboard(storeId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, zap.Error(err))
    }

    return c.JSON(http.StatusOK, categories)
}


func (h *Handler) DeleteCategory(c echo.Context) (error) {
	logger.Debug("Deleting billboard by id...")
	billboardId := c.Param("billboardId")
	storeId := c.Param("storeId")
	userID, ok := c.Get("userID").(string)
	if !ok || userID == "" {
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }
	if !h.storeService.IsOwnerOfStore(storeId, userID) { 
        return c.JSON(http.StatusForbidden, "You are not authorized to update this billboard")
    }
	err := h.billboardService.DeleteBillboard(billboardId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h *Handler) AddCategory(c echo.Context) error {
    logger.Debug("Adding new billboard...")

    storeId := c.Param("storeId")
	userID, ok := c.Get("userID").(string)
	if !ok || userID == "" {
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }
    var req CreateBillboardRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

	if !h.storeService.IsOwnerOfStore(storeId, userID) { 
        return c.JSON(http.StatusForbidden, "You are not authorized to update this billboard")
    }

    billboard := &entity.Billboard{
        Id:        uuid.New().String(),
        StoreId:   storeId,
        Label:     req.Label,
        ImageUrl:  req.ImageUrl,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
        IsActive:  req.IsActive,
    }

    err := h.billboardService.CreateBillboard(billboard)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to create billboard")
    }

    return c.NoContent(http.StatusOK)
}

func (h *Handler) UpdateCategory(c echo.Context) error {
    logger.Debug("Updating billboard...")

    billboardId := c.Param("billboardId")
    userID, ok := c.Get("userID").(string)
    if !ok || userID == "" {
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }

    var req UpdateBillboardRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    logger.Debug("This is the request: ", zap.Reflect("req", req))

    billboard, err := h.billboardService.GetBillboard(billboardId)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, "Billboard not found")
        }
        return c.JSON(http.StatusInternalServerError, "Failed to fetch billboard")
    }

    if !h.storeService.IsOwnerOfStore(billboard.StoreId, userID) { 
        return c.JSON(http.StatusForbidden, "You are not authorized to update this billboard")
    }

    billboard.Label = req.Label
    billboard.ImageUrl = req.ImageUrl
    billboard.UpdatedAt = time.Now()
    billboard.IsActive = req.IsActive

    if err := h.billboardService.UpdateBillboard(billboard); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.NoContent(http.StatusOK)
}