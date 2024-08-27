package handler

import (
	"back-end/internal/domain/entity"
	"back-end/internal/infrastructure/logger"
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"github.com/google/uuid"
)

type CreateBillboardRequest struct {
    StoreId   string `json:"storeId"`
    Label     string `json:"label"`
    ImageUrl  string `json:"imageUrl"`
    IsActive  bool   `json:"isActive"`
}

func (h* Handler) GetBillboardById(c echo.Context) (error) {
	logger.Debug("Fetching billboard by id...")
    billboardId := c.Param("billboardId")
    billboard, err := h.billboardRepo.FindByID(billboardId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    logger.Debug("got billboard", zap.Reflect("billboard", billboard))
    return c.JSON(http.StatusOK, billboard)
}

func (h *Handler) GetBillboardsByStoreId(c echo.Context) (error) {
	logger.Debug("Fetching all billboards by storeId...")
	storeId := c.Param("storeId")
	store, err := h.billboardRepo.GetBillboardsByStoreId(storeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	logger.Debug("got all billboard", zap.Reflect("billboards", store))
	return c.JSON(http.StatusOK, store)
}

func (h *Handler) GetActiveBillboardForSpecificStore(c echo.Context) (error) {
	logger.Debug("Fetching active billboard for a specific store...")
	storeId := c.Param("storeId")
	store, err := h.billboardRepo.GetActiveBillboard(storeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	logger.Debug("got all billboard", zap.Reflect("billboards", store))
	return c.JSON(http.StatusOK, store)
}


func (h *Handler) DeleteByBillboardId(c echo.Context) (error) {
	logger.Debug("Deleting billboard by id...")
	billboardId := c.Param("billboardId")
	err := h.billboardRepo.Delete(billboardId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h *Handler) AddBillboard(c echo.Context) error {
    logger.Debug("Adding new billboard...")

    storeId := c.Param("storeId")
    var req CreateBillboardRequest
    if err := c.Bind(&req); err != nil {
        logger.Error("Failed to bind request: ", zap.Reflect("error", err))
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }
	logger.Debug("this is the request : ",zap.Reflect("req",req))

    billboard := &entity.Billboard{
        Id:        uuid.New().String(),
        StoreId:   storeId,
        Label:     req.Label,
        ImageUrl:  req.ImageUrl,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
        IsActive:  req.IsActive,
    }

	logger.Debug("this is the billboard: ",zap.Reflect("billboard",billboard))

    err := h.billboardRepo.Create(billboard)
    if err != nil {
        logger.Error("Failed to create billboard: ", zap.Reflect("error", err))
        return c.JSON(http.StatusInternalServerError, "Failed to create billboard")
    }

    return c.NoContent(http.StatusOK)
}