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

type ColorRequest struct {
    Name     string `json:"name"`   
    Value  string `json:"value"`  
}

func (h* Handler) GetColorById(c echo.Context) (error) {
	logger.Debug("Fetching color by id...")
    colorId := c.Param("colorId")
    color, err := h.colorService.GetColor(colorId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, color)
}

func (h *Handler) GetAllColors(c echo.Context) error {
    logger.Debug("Fetch All colors...")
    colors, err := h.colorService.GetAllColors()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, zap.Error(err))
    }
    return c.JSON(http.StatusOK, colors)
}


func (h *Handler) DeleteColor(c echo.Context) (error) {
	logger.Debug("Deleting color by id...")
	colorId := c.Param("colorId")
	storeId := c.Param("storeId")
	userID, ok := c.Get("userID").(string)
	if !ok || userID == "" {
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }
	if !h.storeService.IsOwnerOfStore(storeId, userID) { 
        return c.JSON(http.StatusForbidden, "You are not authorized to delete this color")
    }
	err := h.colorService.DeleteColor(colorId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h *Handler) AddColor(c echo.Context) error {
    logger.Debug("Adding new color...")

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
        return c.JSON(http.StatusForbidden, "You are not authorized to update this color")
    }

    color := &entity.Color{
        Id:        uuid.New().String(),
        StoreId:   storeId,
        Name:      req.Name,
        Value:     req.Value,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    err := h.colorService.CreateColor(color)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to create size")
    }

    return c.NoContent(http.StatusOK)
}

func (h *Handler) UpdateColor(c echo.Context) error {
    logger.Debug("Updating color...")

    colorId := c.Param("colorId")
    userID, ok := c.Get("userID").(string)
    if !ok || userID == "" {
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }

    var req ColorRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    color, err := h.colorService.GetColor(colorId)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, "Color not found")
        }
        return c.JSON(http.StatusInternalServerError, "Failed to fetch color")
    }

    if !h.storeService.IsOwnerOfStore(color.StoreId, userID) { 
        return c.JSON(http.StatusForbidden, "You are not authorized to update this color")
    }

    color.Name = req.Name
    color.Value = req.Value
    color.UpdatedAt = time.Now()

    if err := h.colorService.UpdateColor(color); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.NoContent(http.StatusOK)
}