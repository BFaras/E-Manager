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

type SizeRequest struct {
    Name     string `json:"name"`   
    Value  string `json:"value"`  
}

func (h* Handler) GetSizeById(c echo.Context) (error) {
	logger.Debug("Fetching size by id...")
    sizeId := c.Param("sizeId")
    size, err := h.sizeService.GetSize(sizeId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, size)
}

func (h *Handler) GetAllSizes(c echo.Context) error {
    logger.Debug("Fetch sizes with billboard by storeId...")
    sizes, err := h.sizeService.GetAllSizes()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, zap.Error(err))
    }
    return c.JSON(http.StatusOK, sizes)
}


func (h *Handler) DeleteSize(c echo.Context) (error) {
	logger.Debug("Deleting size by id...")
	sizeId := c.Param("sizeId")
	storeId := c.Param("storeId")
	userID, ok := c.Get("userID").(string)
	if !ok || userID == "" {
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }
	if !h.storeService.IsOwnerOfStore(storeId, userID) { 
        return c.JSON(http.StatusForbidden, "You are not authorized to delete this category")
    }
	err := h.sizeService.DeleteSize(sizeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h *Handler) AddSize(c echo.Context) error {
    logger.Debug("Adding new size...")

    storeId := c.Param("storeId")
	userID, ok := c.Get("userID").(string)
	if !ok || userID == "" {
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }
    var req SizeRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

	if !h.storeService.IsOwnerOfStore(storeId, userID) { 
        return c.JSON(http.StatusForbidden, "You are not authorized to update this size")
    }

    category := &entity.Size{
        Id:        uuid.New().String(),
        StoreId:   storeId,
        Name:      req.Name,
        Value:     req.Value,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    err := h.sizeService.CreateSize(category)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to create size")
    }

    return c.NoContent(http.StatusOK)
}

func (h *Handler) UpdateSize(c echo.Context) error {
    logger.Debug("Updating size...")

    sizeId := c.Param("sizeId")
    userID, ok := c.Get("userID").(string)
    if !ok || userID == "" {
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }

    var req SizeRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    size, err := h.sizeService.GetSize(sizeId)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, "Size not found")
        }
        return c.JSON(http.StatusInternalServerError, "Failed to fetch size")
    }

    if !h.storeService.IsOwnerOfStore(size.StoreId, userID) { 
        return c.JSON(http.StatusForbidden, "You are not authorized to update this size")
    }

    size.Name = req.Name
    size.Value = req.Value
    size.UpdatedAt = time.Now()

    if err := h.sizeService.UpdateSize(size); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.NoContent(http.StatusOK)
}