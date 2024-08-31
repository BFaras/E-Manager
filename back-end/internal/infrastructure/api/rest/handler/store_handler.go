package handler

import (
	"back-end/internal/infrastructure/logger"
	"net/http"
	"github.com/labstack/echo/v4"
    "back-end/internal/domain/entity"
    "time"
    "github.com/google/uuid"
    "database/sql"
)

func (h *Handler) GetAllStores(c echo.Context) error {
    logger.Debug("Fetching all stores...")
    store, err := h.storeService.GetAllStores()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, store)
}

func (h *Handler) GetStoreById(c echo.Context) error {
    logger.Debug("Fetching store by storeId...")
    id := c.Param("id")
    store, err := h.storeService.GetStore(id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, store)
}

func (h *Handler) GetStoreByUserId(c echo.Context) error {
    logger.Debug("Fetching store by userId...")
    userId := c.Param("userId")
    store, err := h.storeService.GetByUserId(userId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, store)
}

func (h *Handler) GetStoresByUserId(c echo.Context) error {
    logger.Debug("Fetching all stores by userId...")
    userId := c.Param("userId")
    store, err := h.storeService.GetAllByUserId(userId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, store)
}

func (h *Handler) GetStoreByIdAndUserId(c echo.Context) error {
    logger.Debug("Fetching store by id and userId...")
    userId := c.Param("userId")
    id := c.Param("storeId")
    store, err := h.storeService.GetByIdAndUserId(id,userId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, store)
}

func (h *Handler) DeleteStore(c echo.Context) (error) {
	logger.Debug("Deleting billboard by id...")
	billboardId := c.Param("billboardId")
	storeId := c.Param("storeId")
	userID, ok := c.Get("userID").(string)
	if !ok || userID == "" {
        logger.Error("User ID missing or invalid")
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

func (h *Handler) AddStore(c echo.Context) error {
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

func (h *Handler) UpdateStore(c echo.Context) error {
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

