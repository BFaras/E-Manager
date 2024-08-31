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

type CreateBillboardRequest struct {
    StoreId   string `json:"storeId"`
    Label     string `json:"label"`
    ImageUrl  string `json:"imageUrl"`
    IsActive  bool   `json:"isActive"`
}

type UpdateBillboardRequest struct {
    Label     string `json:"label"`   
    ImageUrl  string `json:"imageUrl"`  
    IsActive  bool   `json:"isActive"`  
}

func (h* Handler) GetBillboardById(c echo.Context) (error) {
	logger.Debug("Fetching billboard by id...")
    billboardId := c.Param("billboardId")
    billboard, err := h.billboardService.GetBillboard(billboardId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, billboard)
}

func (h *Handler) GetBillboardsByStoreId(c echo.Context) (error) {
	logger.Debug("Fetching all billboards by storeId...")
	storeId := c.Param("storeId")
	store, err := h.billboardService.GetBillboardsByStoreId(storeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, store)
}

func (h *Handler) GetActiveBillboardForSpecificStore(c echo.Context) (error) {
	logger.Debug("Fetching active billboard for a specific store...")
	storeId := c.Param("storeId")
	store, err := h.billboardService.GetActiveBillboard(storeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	logger.Debug("got all billboard", zap.Reflect("billboards", store))
	return c.JSON(http.StatusOK, store)
}


func (h *Handler) DeleteBillboard(c echo.Context) (error) {
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

func (h *Handler) AddBillboard(c echo.Context) error {
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
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.NoContent(http.StatusOK)
}

func (h *Handler) UpdateBillboard(c echo.Context) error {
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

    billboard, err := h.billboardService.GetBillboard(billboardId)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, err.Error())
        }
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    if !h.storeService.IsOwnerOfStore(billboard.StoreId, userID) { 
        return c.JSON(http.StatusForbidden, err.Error())
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