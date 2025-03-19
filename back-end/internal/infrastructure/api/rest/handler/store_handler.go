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

type StoreRequest struct {
    Name     string `json:"name"`   
}

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
	logger.Debug("Deleting store by id...")
	storeId := c.Param("storeId")
	userID, ok := c.Get("userID").(string)
	if !ok || userID == "" {
        logger.Error("User ID missing or invalid")
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }
	if !h.storeService.IsOwnerOfStore(storeId, userID) { 
        return c.JSON(http.StatusForbidden, "You are not authorized to update this billboard")
    }
	err := h.storeService.DeleteStore(storeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h *Handler) AddStore(c echo.Context) error {
    logger.Debug("Adding new store...")

	userID, ok := c.Get("userID").(string)
	if !ok || userID == "" {
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }

    var req StoreRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    store := &entity.Store{
        Id:        uuid.New().String(),
        Name:       req.Name,
        UserId:   userID,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    err := h.storeService.CreateStore(store)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    logger.Debug("New store added:", zap.Any("store",store))

    return c.JSON(http.StatusCreated,store)
}

func (h *Handler) UpdateStore(c echo.Context) error {
    logger.Debug("Updating store...")

    storeId := c.Param("storeId")
    userID, ok := c.Get("userID").(string)
    if !ok || userID == "" {
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }

    var req StoreRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    store, err := h.storeService.GetStore(storeId)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, err.Error())
        }
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    if !h.storeService.IsOwnerOfStore(store.Id, userID) { 
        return c.JSON(http.StatusForbidden, "You are not authorized to update this billboard")
    }

    store.Name = req.Name
    store.UpdatedAt = time.Now()

    if err := h.storeService.UpdateStore(store); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.NoContent(http.StatusOK)
}

